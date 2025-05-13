package components

import (
	"syscall/js"
	"ch/oliumbi/halogen/events"
)

type Component struct {
	reference js.Value
	children []*Component
	observers []*events.Observer
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

func (component *Component) Unmount() {
	component.unmountChildren()

	for _, observer := range component.observers {
		events.Unsubscribe(observer)
	}

	component.reference.Call("remove")
}

func (component *Component) unmountChildren() {
	for _, child := range component.children {
		child.Unmount()
	}
}

func (component *Component) SetText(text string) {
	component.unmountChildren()

	component.reference.Set("innerText", text)
}

func (component *Component) Child(child *Component) {
	component.children = append(component.children, child)

	component.reference.Call("appendChild", child.reference)
}

func (component *Component) Listener(event string, callback func(args []js.Value)) {
	component.reference.Call("addEventListener", event, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		callback(args)
		return nil
	}))
}

func (component *Component) Event(name string, callback func(payload any)) {
	var observer = &events.Observer{
		Name: name,
		Callback: callback,
	}

	component.observers = append(component.observers, observer)

	events.Subscribe(observer)
}

func (component *Component) Click(callback func()) {
	component.Listener("click", func(args []js.Value) {
		callback()
	})
}

func (component *Component) Change(callback func(value string)) {
	component.Listener("change", func(args []js.Value) {
		callback(args[0].Get("target").Get("value").String())
	})
}
