package gui

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

// ScreenOne is
func ScreenOne() {
	application := app.New()

	w := application.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Siv..!"),
		widget.NewButton("Quit", func() {
			application.Quit()
		}),
	))

	w.ShowAndRun()
}
