package views

import (
	"ch/oliumbi/halogen/components"
)

func Home() {

	root := components.Root()

	container := components.Create("div")

	text := components.Create("p")
	text.SetText("hello")

	button := components.Create("button")
	button.SetText("click")
	button.OnClick(func() {
		text.SetText("bye")
		container.Remove()
	})

	button2 := components.Create("button")
	button2.SetText("click")
	button2.OnClick(func() {
		text.SetText("yeet")
	})

	input := components.Create("input")
	input.OnChange(func(value string) {
		text.SetText(value)
	})

	root.AppendChild(container)
	root.AppendChild(button2)
	root.AppendChild(input)

	container.AppendChild(text)
	container.AppendChild(button)
}
