package main

import (
	"encoding/hex"
	"log"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Creat the widgets that are needed for the OPc calculation form
// Save the Entry widgets to the global cfg as we will need to access them
// from other functions and we will need them for testing
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

// Create the OPc Form button bar
func (app *config) createOPCButtons() *fyne.Container {
	calc := widget.NewButton("Calculate", app.calcOPCFunc())
	clear := widget.NewButton("Clear", app.clearOPCFunc())
	sp := layout.NewSpacer()

	con := container.New(layout.NewHBoxLayout(), sp, clear, calc)

	return con
}

//Implement the button functions as seperate functions returning a function
//this is just to kep the code clearer and more readable

// The calc function for OPC Form
func (app *config) calcOPCFunc() func() {
	return func() {
		log.Println("Ki:", app.KiEntry.Text)
		log.Println("OP:", app.OPEntry.Text)
		app.OPcVal.SetText(calculateOPC(app.KiEntry.Text, app.OPEntry.Text))
	}
}

// The clear function for OPC Form
func (app *config) clearOPCFunc() func() {
	return func() {
		app.KiEntry.SetText("")
		app.OPEntry.SetText("")
		app.OPcVal.SetText("")
	}
}

// Create the coopy button and layout
// use to copy the OPc value when calcualted
// Call it tool bar as we may add more tool functions at a later date
func (app *config) createCopyBar() *fyne.Container {
	copy_tool := widget.NewButtonWithIcon("Copy OPc", theme.ContentCopyIcon(), func() {
		app.mainWindow.Clipboard().SetContent(app.OPcVal.Text)
	})

	//generate a random Ki value
	rnd_ki := widget.NewButtonWithIcon("Rnd Ki", theme.ContentPasteIcon(), func() {
		rnd_src := rand.New(rand.NewSource(time.Now().UnixNano()))
		b := make([]byte, 16)
		if _, err := rnd_src.Read(b); err != nil {
			panic(err)
		}
		cfg.KiEntry.Text = hex.EncodeToString(b)
		cfg.KiEntry.Refresh()
	})

	tool_bar := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), rnd_ki, copy_tool)

	return tool_bar
}

func (app *config) createMainUI() *fyne.Container {

	sep := widget.NewSeparator()
	opcf := cfg.createOPCForm()
	opcb := cfg.createOPCButtons()
	cb := cfg.createCopyBar()

	//pull the container out so that it can be referenced in the TabItem
	opc_form := container.New(layout.NewVBoxLayout(), sep, opcf, opcb, cb)
	//form for the K4 calculations
	k4form := app.createK4UI()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("OPc Calculation", theme.HomeIcon(), opc_form),
		container.NewTabItemWithIcon("K4 Calculation", theme.MediaPlayIcon(), k4form),
	)
	// return the following container
	return container.New(layout.NewPaddedLayout(), tabs)
}
