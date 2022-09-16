package gx

import (
	"fmt"
	"syscall/js"
)

func Render(c *ComponentInstance, anchor string) {
	fmt.Println("GX: Rendering...")
	ch := make(chan struct{}, 0)
	app := js.Global().Get("document").Call("getElementById", anchor)
	fmt.Println("GX: App:", app.Get("id").String())

	c.ctx = c.i(c, nil, c.invalidate)

	_c := c.c(c.ctx)
	c.self = &_c
	c.self.C()
	c.self.M(app, js.Null())
	// c.D()
	<-ch
}

func (c *ComponentInstance) stale(i int) bool {
	return (c.vars[0] & uint64(i)) != 0
}

func (c *ComponentInstance) invalidate(i int, v any) {
	c.ctx[i] = v
	c.vars[0] |= 1 << uint64(i)
	c.self.U(c.ctx, c.stale)
}

func NewComponent(i func(*ComponentInstance, *map[string]any, InvalidateFn) []any, c func([]any) Component) *ComponentInstance {
	return &ComponentInstance{
		vars: []uint64{0b0000000000000000000000000000000000000000000000000000000000000000},
		ctx:  make([]any, 0),
		i:    i,
		c:    c,
	}
}

func Element(tag string) *El {
	return &El{
		Tag:      tag,
		Props:    make(map[string]any),
		Children: make([]El, 0),
	}
}

func Text(text any) *TextEl {
	return &TextEl{
		content: fmt.Sprintf("%v", text),
	}
}

func (t *TextEl) Create() js.Value {
	e := js.Global().Get("document").Call("createTextNode", t.content)
	t.self = e
	return e
}

func (t *TextEl) Ref() js.Value {
	return t.self
}

func (t *TextEl) Remove() {
	t.self.Get("parentNode").Call("removeChild", t.self)
}

func (el *El) Create() js.Value {
	e := js.Global().Get("document").Call("createElement", el.Tag)
	for k, v := range el.Props {
		e.Call("setAttribute", k, fmt.Sprintf("%v", v))
	}

	el.self = e
	return e
}

func (el *El) Ref() js.Value {
	return el.self
}

func (el *El) Remove() {
	el.self.Get("parentNode").Call("removeChild", el.self)
}

func Insert(target js.Value, el Node, anchor js.Value) {
	target.Call("insertBefore", el.Create(), anchor)
}

func Append(target Node, el Node) {
	target.Ref().Call("appendChild", el.Create())
}

func Detach(target Node) {
	target.Remove()
}

func Attr(target *El, key string, value any) {
	target.Props[key] = value
}

func Listen(target Node, event string, handler any) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler.(func([]js.Value))(args)
		return nil
	})
	target.Ref().Call("addEventListener", event, cb)
}

func Set(target Node, content any) {
	target.Ref().Set("textContent", fmt.Sprintf("%v", content))
}
