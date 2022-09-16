package gx

import "syscall/js"

type InvalidateFn func(i int, v any)
type StaleFn func(i int) bool

type Component struct {
	C func()
	M func(js.Value, js.Value)
	U func([]any, StaleFn)
	D func()
}

type ComponentInstance struct {
	vars []uint64
	ctx  []any
	i    func(*ComponentInstance, *map[string]any, InvalidateFn) []any
	c    func([]any) Component
	self *Component
}

type El struct {
	Tag      string
	Props    map[string]any
	Children []El
	self     js.Value
}

type TextEl struct {
	content string
	self    js.Value
}

type Node interface {
	Create() js.Value
	Ref() js.Value
	Remove()
}
