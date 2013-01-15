package msgui

import (
	W "github.com/lxn/go-winapi"
)

// RectArea is the base struct of window weights.
type RectArea struct {
	OrignX, OrignY, Width, Height int32
}

// ------------------------------------------------------------
type Drawer interface {Draw(W.HDC)}

type Component struct {
	RectArea
	parent, previous, next *Component
	Value Drawer
}

func (c *Component) GetParent() *Component {return parent}
func (c *Component) Privious() *Component {return previous}
func (c *Component) Next() *Component {return next}

type Components struct {
	Component
	head, tail *Component
	length uint32
}

func (cs *Components) Len() uint32 {return cs.lenth}
func (cs *Components) Head() *Component {return cs.head}
func (cs *Components) Tail() *Component {return cs.tail}

func (cs *Components) AddComponent(c *Component) {
	if cs.tail == nil {
		cs.head, cs.tail = c, c
		c.previous, c.next = nil, nil
		cs.length = 1
	} else {
		c.previous, c.next = cs.tail.previous, nil
		cs.tail.next = c
		cs.tail = c
		cs.length++
	}
	c.parent = cs
}

func (cs *Components) DelComponent(c *Component) {
	if c.parent != cs {return}
	if c.previous == nil {
		cs.head = c.next
	} else {
		c.previous.next = c.next
	}
	if c.next == nil {
		cs.tail = c.previous
	} else {
		c.next.previous = c.previous
	}
	c.previous, c.next, c.parent = nil, nil, nil
	cs.length--
}

func (cs *Components) MoveToNext(c *Component) {
	if c.parent != cs || cs.tail == c {return}
	cp, cn, cnn := c.previous, c.next, c.next.next
	c.next, cn.next = cn.next, c
	c.previous, cn.previous = cn, c.previous
	if cp != nil {cp.next = cn}
	if cnn != nil {cnn.previous = c}
}

func (cs *Components) MoveToPrevious(c *Component) {
	if c.parent != cs || cs.head == c {return}
	cn, cp, cpp := c.next, c.previous, c.previous.previous
	c.previous, cp.previous = cp.previous, c
	c.next, cp.next = cp, c.next
	if cn != nil {cn.previous = cp}
	if cpp != nil {cpp.next = c}
}

func (cs *Components) MoveToHead(c *Component) {
	if c.parent != cs || cs.head == c {return}
	if c.next != nil {c.next.previous = c.previous}
	c.previous.next = c.next
	c.next, c.previous, cs.head.previous = cs.head, nil, c
	cs.head = c
}

func (cs *Components) MoveToTail(c *Component) {
	if c.parent != cs || cs.tail == c {return}
	if c.previous != nil {c.previous.next = c.next}
	c.next.previous = c.previous
	c.previous, c.next, cs.tail.next = cs.tail, nil, c
	cs.tail = c
}
