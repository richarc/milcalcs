package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (app *config) createK4UI() *fyne.Container {

	//entry and lable for the K4 key
	k4l := widget.NewLabel("Key Value:")
	k4Entry := widget.NewEntry()

	//options for slecting the encryption cypher
	opts := []string{
		"3DES",
		"XXXX",
		"YYYY",
	}
	//label for cypher options
	optsl := widget.NewLabel("Select Cypher:")
	k4Options := widget.NewSelect(opts, func(string) {})

	//What are we encrypting? label and entry for the value to be encrypted
	ekol := widget.NewLabel("Ki or OPc:")
	ekoEntry := widget.NewEntry()

	//output for the encrypted value
	eoutl := widget.NewLabel("eKi or eOPc:")
	eoutEntry := widget.NewLabel("")

	k4b := cfg.createK4Buttons()

	k4ui := container.New(layout.NewFormLayout(), k4l, k4Entry, optsl, k4Options, ekol, ekoEntry, eoutl, eoutEntry)

	return container.New(layout.NewVBoxLayout(), k4ui, k4b)

}

//Create the k4 buttone
func (app *config) createK4Buttons() *fyne.Container {
	calc := widget.NewButton("Calculate", func() {})
	clear := widget.NewButton("Clear", func() {})
	sp := layout.NewSpacer()

	con := container.New(layout.NewHBoxLayout(), sp, clear, calc)

	return con
}
