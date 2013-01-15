/* This package is used to create gui components.
 * author: musen
 * version: 0.0.1
 * copyright: musen
 * history
     2013-1-13 0.0.1 create by musen
*/
package msgui
import (
	W "github.com/lxn/go-winapi"
	"syscall"
//	"os"
	"unsafe"
)

// hInstance of application which use this package.
var hInstance W.HINSTANCE = W.GetModuleHandle(nil)

// Windows map of application
var winEventMap = make(tWinMap)

// default main window's hwnd
var defaultWindow W.HWND

//contain all window's hwnd and component of application.
var windows = make(map[W.HWND]interface{})

type tWinMap map[W.HWND]map[uint32]TWinProc

// deal with procdure of special message of window using message identity
type TWinProc func(W.HWND, uintptr, uintptr) uintptr

// Create map of the window by using hwnd.
func (winMap tWinMap) createMap (hwnd W.HWND) {
	_, ok := winMap[hwnd]
	if !ok {
		winMap[hwnd] = make(map[uint32]TWinProc)
	}
}

// add a procedure of window to tWinMap var.
func (winMap tWinMap) addProc (hwnd W.HWND, msg uint32, proc TWinProc) {
	_, ok := winMap[hwnd]
	if !ok {winMap.createMap(hwnd)}
	winMap[hwnd][msg] = proc
}

func (winMap tWinMap) dispatch (hwnd W.HWND, msg uint32, wParam, lParam uintptr) {
	_, ok := winMap[hwnd]
	if ok {
		f, ok := winMap[hwnd][msg]
		if ok {
			f(hwnd, wParam, lParam)
		}
		if msg == W.WM_DESTROY {
			quitMsgProc(hwnd)
		}
	}
}

func quitMsgProc(hwnd W.HWND) {
	if hwnd == defaultWindow {
		W.PostQuitMessage(0)
	} else {
		W.ShowWindow(hwnd, W.SW_MINIMIZE)
	}
}

func wndProc(hwnd W.HWND, msg uint32, wparam, lparam uintptr) uintptr {
	winEventMap.dispatch(hwnd, msg, wparam, lparam)
	return W.DefWindowProc(hwnd, msg, wparam, lparam)
}

/* Regist a window's class use winapi.the name is the name of window's class.
 * wndProc is the procedure of this window.
 */
func registFrameClass(name *uint16, wndProc uintptr) W.ATOM {
	wc := new (W.WNDCLASSEX)
	wc.CbSize = uint32(unsafe.Sizeof(*wc))
	wc.Style = W.CS_HREDRAW | W.CS_VREDRAW
	wc.LpfnWndProc = wndProc
	wc.CbClsExtra = 0
	wc.CbWndExtra = 0
	wc.HInstance = hInstance
	wc.HIcon = 0
	wc.HCursor = W.LoadCursor(0, W.MAKEINTRESOURCE(W.IDC_ARROW))
	wc.HbrBackground = W.HBRUSH(W.WHITE_BRUSH)
	wc.LpszMenuName = nil
	wc.LpszClassName = name
	wc.HIconSm = 0
	return W.RegisterClassEx(wc)
}

// RectArea is the base struct of window weights.
type RectArea struct {
	OrignX, OrignY, Width, Height int32
}

// Create a window's window,which has a hwnd handle.
func CreateWindow(title, name string, rect RectArea) (hwnd W.HWND) {
	var namePtr *uint16 = W.StringToBSTR(name)
	registFrameClass(namePtr, syscall.NewCallback(wndProc))
	hwnd = W.CreateWindowEx(
		W.WS_EX_CLIENTEDGE,
		namePtr,
		W.StringToBSTR(title),
		W.WS_OVERLAPPEDWINDOW,
		rect.OrignX,
		rect.OrignY,
		rect.Width,
		rect.Height,
		0,
		0,
		hInstance,
		unsafe.Pointer(nil))
	if defaultWindow == 0 {SetDefault(hwnd)}
	winEventMap.createMap(hwnd)
	windows[hwnd] = title
	return
}

// set default window of application
func SetDefault(hwnd W.HWND) {
	defaultWindow = hwnd
}

// Start a gui application
func Start(hwnd W.HWND) uintptr {
	var message W.MSG
//	W.ShowWindow(hwnd,W.SW_SHOW)
	for{
		if W.GetMessage(&message, 0, 0, 0) == 0 {break}
		W.TranslateMessage(&message)
		W.DispatchMessage(&message)
	}
	return 0
}
	