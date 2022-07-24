package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// create a type that allows us to access Ui elements across scope of different functions
type config struct {
	KiEntry    *widget.Entry
	OPEntry    *widget.Entry
	OPcVal     *widget.Label
	mainWindow fyne.Window
}

//Declare as a package global type
var cfg config

func main() {
	a := app.New()
	cfg.mainWindow = a.NewWindow("Milenage calculations")

	mainUI := cfg.createMainUI()
	cfg.mainWindow.SetContent(mainUI)

	cfg.mainWindow.Resize(fyne.Size{Width: 600, Height: 500})
	cfg.mainWindow.ShowAndRun()
}
