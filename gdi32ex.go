package msgui

import (
	W "github.com/lxn/go-winapi"
	"syscall"
)

var (
	// Library
	libgdi32 uintptr
	// Functions
	roundRect uintptr
)

func init() {
	// Library
	libgdi32 = W.MustLoadLibrary("gdi32.dll")
	// Functions
	roundRect = W.MustGetProcAddress(libgdi32, "RoundRect")
}

func RoundRect(hdc W.HDC, nLeftRect, nTopRect, nRightRect, nBottomRect, nWidth, nHeight int32) bool {
	ret, _, _ := syscall.Syscall9(roundRect, 7,
		uintptr(hdc),
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect),
		uintptr(nWidth),
		uintptr(nHeight),
		0, 0)

	return ret != 0
}

