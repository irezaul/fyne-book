package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myapp := app.New()
	window := myapp.NewWindow("New CRP")
	window.Resize(fyne.NewSize(600, 400))

	// create a titles

	title := canvas.NewText("login page", color.Opaque)
	title.TextSize = 20
	title.Alignment = fyne.TextAlignCenter

	status := widget.NewLabel("Welcome to login page !")

	// twoForm added in one form
	form := widget.NewForm(

		widget.NewFormItem("UserName", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)
	window.SetContent(
		container.NewVBox(
			title,
			form,
			status,
			widget.NewButton("Login... :)", func() {
				status.SetText("logine successfully :)")

			}),
		))

	window.ShowAndRun()
}
