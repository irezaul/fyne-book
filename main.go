package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//new app
	app := app.New()

	// new window
	window := app.NewWindow("miniERP v0.0.1")

	// window resize
	window.Resize(fyne.NewSize(600, 400))

	// Window app Icon set
	window.SetIcon(theme.FyneLogo())

	// cusotme icon
	icon, _ := fyne.LoadResourceFromPath("./assets/img/master.png")

	//=============================================================//
	//content start //
	mainTile := canvas.NewText("miniERP demo", color.White) // new tite text & color set
	mainTile.TextSize = 24                                  // font size of title
	mainTile.Alignment = fyne.TextAlignCenter               // title alignment set

	// input 1 email //
	email := widget.NewEntry()
	email.SetPlaceHolder("Enter Your EMAIL")
	// end //

	// input 2 password//
	password := widget.NewEntry()
	password.SetPlaceHolder("Password Required")
	// end 2 //

	//save button//
	save_btn := widget.NewButton("save", func() {

	})

	//content set
	window.SetIcon(icon)
	window.SetContent(container.NewVBox(

		mainTile,
		email,
		password,
		save_btn,
	))

	// create & run
	window.ShowAndRun()

}
