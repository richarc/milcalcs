package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (app *config) createMilUi() *fyne.Container {

	//Need to find a better way to create this form

	//labels and values for the r constants
	r1l := widget.NewLabel("r1:")
	r1v := widget.NewLabel(fmt.Sprintf("Hex value = ( %#x ) Decimal value = ( %d )", milvals.r1, milvals.r1))

	r2l := widget.NewLabel("r2:")
	r2v := widget.NewLabel(fmt.Sprintf("Hex value = ( %#x ) Decimal value = ( %d )", milvals.r2, milvals.r2))

	r3l := widget.NewLabel("r3:")
	r3v := widget.NewLabel(fmt.Sprintf("Hex value = ( %#x ) Decimal value = ( %d )", milvals.r3, milvals.r3))

	r4l := widget.NewLabel("r4:")
	r4v := widget.NewLabel(fmt.Sprintf("Hex value = ( %#x ) Decimal value = ( %d )", milvals.r4, milvals.r4))

	r5l := widget.NewLabel("r4:")
	r5v := widget.NewLabel(fmt.Sprintf("Hex value = ( %#x ) Decimal value = ( %d )", milvals.r5, milvals.r5))

	//the labels and values for the c constants

	c1l := widget.NewLabel("c1:")
	c1v := widget.NewLabel(fmt.Sprintf("Hex value = ( %#x )", milvals.c1))

	c2l := widget.NewLabel("c2:")
	c2v := widget.NewLabel(fmt.Sprintf("Hex value = ( %#x )", milvals.c2))

	c3l := widget.NewLabel("c3:")
	c3v := widget.NewLabel(fmt.Sprintf("Hex value = ( %#x )", milvals.c3))

	c4l := widget.NewLabel("c4:")
	c4v := widget.NewLabel(fmt.Sprintf("Hex value = ( %#x )", milvals.c4))

	c5l := widget.NewLabel("c5:")
	c5v := widget.NewLabel(fmt.Sprintf("Hex value = ( %#x )", milvals.c5))

	return container.New(layout.NewFormLayout(), r1l, r1v, r2l, r2v, r3l, r3v, r4l, r4v, r5l, r5v,
		c1l, c1v, c2l, c2v, c3l, c3v, c4l, c4v, c5l, c5v)
}
