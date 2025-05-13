package views

import (
	"ch/oliumbi/halogen/components"
)

func Home() {

	root := components.Root()

	container := components.Create("div")
	text := components.Create("p")
	button := components.Create("button")
	button2 := components.Create("button")
	input := components.Create("input")

	text.SetText("hello")

	button.SetText("click")
	button.Click(func() {
		text.SetText("bye")
		container.Unmount()
	})

	button2.SetText("click")
	button2.Click(func() {
		text.SetText("yeet")
	})

	input.Change(func(value string) {
		text.SetText(value)
	})

	container.Child(text)
	container.Child(button)

	root.Child(container)
	root.Child(button2)
	root.Child(input)
}
