package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (app *config) createMilUi() *fyne.Container {

	var data = [][]string{
		{"r1:", fmt.Sprintf("%v", (milvals.r1))},
		{"r2:", fmt.Sprintf("%v", (milvals.r2))},
		{"r3:", fmt.Sprintf("%v", (milvals.r3))},
		{"r4:", fmt.Sprintf("%v", (milvals.r4))},
		{"r5:", fmt.Sprintf("%v", (milvals.r5))},
		{"128 bit", "Only last 32 bits are shown"},
		{"c1:", fmt.Sprintf("%032b", (milvals.c1))},
		{"c2:", fmt.Sprintf("%032b", (milvals.c2))},
		{"c3:", fmt.Sprintf("%032b", (milvals.c3))},
		{"c4:", fmt.Sprintf("%032b", (milvals.c4))},
		{"c5:", fmt.Sprintf("%032b", (milvals.c5))},
	}

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	return container.New(layout.NewPaddedLayout(), list)
}
