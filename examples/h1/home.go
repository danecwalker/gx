package main

import (
  "syscall/js"
  . "github.com/danecwalker/gx/pkg/gxruntime"
)

var (
  v = []uint64{ 0b0000000000000000000000000000000000000000000000000000000000000000 }
)
func instance(self *ComponentInstance, _props *map[string]any, _invalidate InvalidateFn) []any {
  return []any{}
}
func create_component(ctx []any) Component {
  var h10 *El
  var t0 *TextEl
  var h20 *El
  var t1 *TextEl
  var h30 *El
  var t2 *TextEl
  var h40 *El
  var t3 *TextEl
  var h50 *El
  var t4 *TextEl
  var h60 *El
  var t5 *TextEl
  var mounted bool
  return Component{
    C: func() {
      h10 = Element("h1")
      t0 = Text("My Awesome Nav in GX")
      h20 = Element("h2")
      t1 = Text("My Awesome Nav in GX")
      h30 = Element("h3")
      t2 = Text("My Awesome Nav in GX")
      h40 = Element("h4")
      t3 = Text("My Awesome Nav in GX")
      h50 = Element("h5")
      t4 = Text("My Awesome Nav in GX")
      h60 = Element("h6")
      t5 = Text("My Awesome Nav in GX")
    },
    M: func(target js.Value, anchor js.Value) {
      Insert(target, h10, anchor)
      Insert(target, h20, anchor)
      Insert(target, h30, anchor)
      Insert(target, h40, anchor)
      Insert(target, h50, anchor)
      Insert(target, h60, anchor)
      Append(h10, t0)
      Append(h20, t1)
      Append(h30, t2)
      Append(h40, t3)
      Append(h50, t4)
      Append(h60, t5)
      if !mounted { mounted = true }
    },
    D: func() {
      Detach(h10)
      Detach(t0)
      Detach(h20)
      Detach(t1)
      Detach(h30)
      Detach(t2)
      Detach(h40)
      Detach(t3)
      Detach(h50)
      Detach(t4)
      Detach(h60)
      Detach(t5)
      mounted = false
    },
  }
}



func Home() *ComponentInstance {
	println("goodbye")
	return NewComponent(instance, create_component)
}
