package components

import (
	"syscall/js"
)

type Component struct {
	reference js.Value
}

func Root() *Component {
	return &Component{
		reference: js.Global().Get("document").Get("body"),
	}
}

func Create(tag string) *Component {
	return &Component{
		reference: js.Global().Get("document").Call("createElement", tag),
	}
}

func (component *Component) Remove() {
	component.reference.Call("remove")
}

func (component *Component) SetText(text string) {
	component.reference.Set("innerText", text)
}

func (component *Component) OnClick(callback func()) {
	component.reference.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		callback()
		return nil
	}))
}

func (component *Component) OnChange(callback func(value string)) {
	component.reference.Set("onchange", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		callback(args[0].Get("target").Get("value").String())
		return nil
	}))
}

func (component *Component) AppendChild(child *Component) {
	component.reference.Call("appendChild", child.reference)
}
