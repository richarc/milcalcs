package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// create a type that allows us to access Ui elements across scope of different functions
type config struct {
	KiEntry *widget.Entry
	OPEntry *widget.Entry
	OPcVal  *widget.Label
}

//Declare as a global type
var cfg config

//Creat the widgets that are needed for the OPc calculation form
//Save the Entry widgets to the global cfg as we will need to access them
//from other functions and we will need them for testing
func (app *config) createOPCForm() *fyne.Container {
	ki := widget.NewEntry()
	op := widget.NewEntry()
	opc := widget.NewLabel("")

	//Save the Entry widgets to the cfg
	app.KiEntry = ki
	app.OPEntry = op
	app.OPcVal = opc

	//We don't need to save the label widgets as no need to access later
	kil := widget.NewLabel("Ki Value:")
	opl := widget.NewLabel("OP Value:")
	opcl := widget.NewLabel("OPc Result:")

	con := container.New(layout.NewFormLayout(), kil, ki, opl, op, opcl, opc)
	//Return a container
	return con
}

//Create the OPc Form button bar
func (app *config) createOPCButtons() *fyne.Container {
	calc := widget.NewButton("Calculate", app.calcFunc())
	clear := widget.NewButton("Clear", app.clearFunc())
	sp := layout.NewSpacer()

	con := container.New(layout.NewHBoxLayout(), sp, clear, calc)

	return con
}

//Implement the button functions as seperate functions returning a function
//this is just to kep the code clearer and more readable

//The calc function for OPC Form
func (app *config) calcFunc() func() {
	return func() {
		log.Println("Ki:", app.KiEntry.Text)
		log.Println("OP:", app.OPEntry.Text)
		app.OPcVal.SetText(calculateOPC(app.KiEntry.Text, app.OPEntry.Text))
	}
}

//The clear function for OPC Form
func (app *config) clearFunc() func() {
	return func() {
		app.KiEntry.SetText("")
		app.OPEntry.SetText("")
		app.OPcVal.SetText("")
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello Window")

	opcf := cfg.createOPCForm()
	opcb := cfg.createOPCButtons()

	w.SetContent(container.New(layout.NewVBoxLayout(), opcf, opcb))

	w.Resize(fyne.Size{Width: 500, Height: 500})
	w.ShowAndRun()
}
