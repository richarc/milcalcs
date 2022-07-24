package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

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

//Create the coopy button and layout
//use to copy the OPc value when calcualted
//Call it tool bar as we may add more tool functions at a later date
func (app *config) createCopyBar(win fyne.Window) *fyne.Container {
	copy_tool := widget.NewButtonWithIcon("Copy OPc", theme.ContentCopyIcon(), func() {
		win.Clipboard().SetContent(app.OPcVal.Text)
	})
	tool_bar := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), copy_tool)

	return tool_bar
}

func (app *config) createMainUI() *fyne.Container {

	tb := cfg.createToolBar()
	opcf := cfg.createOPCForm()
	opcb := cfg.createOPCButtons()
	cb := cfg.createCopyBar(cfg.mainWindow)

	//pull the container out so that it can be referenced in the TabItem
	opc_form := container.New(layout.NewVBoxLayout(), opcf, opcb, cb)

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("OPc Calculation", theme.HomeIcon(), opc_form),
		container.NewTabItemWithIcon("K4 Calculation", theme.MediaPlayIcon(), canvas.NewText("K4 Window", nil)),
		container.NewTabItemWithIcon("Milenage R+C", theme.MediaPlayIcon(), canvas.NewText("Milenage R and C Values", nil)),
	)
	// return the following container
	return container.New(layout.NewVBoxLayout(), tb, tabs)
}
