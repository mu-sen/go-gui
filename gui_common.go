package msgui

import (
	W "github.com/lxn/go-winapi"
)

// RectArea is the base struct of window weights.
type RectArea struct {
	OrignX, OrignY, Width, Height int32
}

type Drawer interface {
	Draw(hdc W.HDC)
}

type Container struct {
	ChildArray 