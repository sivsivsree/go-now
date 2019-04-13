package gui


import (
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)

func ScreenOne() {
	application := app.New()

	w := application.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			application.Quit()
		}),
	))

	w.ShowAndRun()
}