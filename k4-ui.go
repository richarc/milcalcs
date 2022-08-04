package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (app *config) createK4UI() *fyne.Container {

	//entry and lable for the K4 key
	k4l := widget.NewLabel("Key Value:")
	k4k := widget.NewEntry()

	//options for slecting the encryption cypher
	opts := []string{
		"DES ECB",
		"3DES ECB",
		"AES 128 ECB",
	}
	//label for cypher options
	optsl := widget.NewLabel("Select Cypher:")
	k4opts := widget.NewSelect(opts, func(string) {})

	//What are we encrypting? label and entry for the value to be encrypted
	ekol := widget.NewLabel("Ki or OPc:")
	ekoe := widget.NewEntry()

	//output for the encrypted value
	eoutl := widget.NewLabel("eKi or eOPc:")
	eoute := widget.NewLabel("")

	//save the entry widgest to the cfg struct
	app.K4KeyEntry = k4k
	app.K4OptEntry = k4opts
	app.EkoEntry = ekoe
	app.EoutEntry = eoute

	k4b := cfg.createK4Buttons()

	k4ui := container.New(layout.NewFormLayout(), k4l, k4k, optsl, k4opts, ekol, ekoe, eoutl, eoute)

	return container.New(layout.NewVBoxLayout(), k4ui, k4b)
}

//Create the k4 buttons
func (app *config) createK4Buttons() *fyne.Container {
	calc := widget.NewButton("Calculate", app.calcK4Func())
	clear := widget.NewButton("Clear", app.clearK4Func())
	sp := layout.NewSpacer()

	con := container.New(layout.NewHBoxLayout(), sp, clear, calc)

	return con
}

//Implement the button functions as seperate functions returning a function
//this is just to kep the code clearer and more readable

//The calc function for K4 Form
func (app *config) calcK4Func() func() {
	return func() {
		log.Println("K4:", app.K4KeyEntry.Text)
		log.Println("Alg:", app.K4OptEntry.Selected)
		log.Println("kio:", app.EkoEntry.Text)
		app.EoutEntry.SetText(calculateK4out(app.K4KeyEntry.Text, app.K4OptEntry.Selected, app.EkoEntry.Text))
	}
}

//The clear function for K4 Form
func (app *config) clearK4Func() func() {
	return func() {
		app.K4KeyEntry.SetText("")
		app.K4OptEntry.ClearSelected()
		app.EkoEntry.SetText("")
		app.EoutEntry.SetText("")
	}
}
