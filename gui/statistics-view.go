package gui

import (
	"fmt"
	"sparta/file"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func specificValues(data []file.Exercise) map[string]float64 {

	// Make the map that contains each exercise and how many times it has been used.
	values := make(map[string]float64)

	// Bump the count on each iteration.
	for i := range data {
		values[data[i].Activity]++
	}

	return values
}

func (u *user) StatisticsView() fyne.CanvasObject {

	// Get the map with exercises.
	mapex := specificValues(u.Data.Exercise)

	// Make the initial chart values populated with the length of the user data.
	values := make([]chart.Value, len(mapex))

	// Create index for setting the values.
	index := 0

	// Loop through the values and add the data.
	for k, v := range mapex {

		values[index] = chart.Value{
			Value: v,
			Label: k,
		}

		index++
	}

	// Set up the pie chart
	pie := chart.PieChart{
		Width:      512,
		Height:     512,
		Values:     values,
		Background: chart.Style{FillColor: drawing.ColorTransparent},
	}

	// Create an image writer.
	writer := &chart.ImageWriter{}

	// Render ther graph.
	pie.Render(chart.PNG, writer)

	// Get the image to display:
	image, err := writer.Image()
	if err != nil {
		fmt.Println(err)
	}

	return fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(512, 512)), canvas.NewImageFromImage(image))
}
