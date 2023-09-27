package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ellypaws/stencil/drawer"
	imageprocessor "github.com/ellypaws/stencil/image-processor"
	"strconv"
)

type rect struct {
	startX, startY, endX, endY int
}

var movements []imageprocessor.Movement

func startSketching(val string) {
	if len(movements) == 0 {
		fmt.Println("Movements not loaded yet.")
		return
	}

	n, _ := strconv.Atoi(val)
	if n > len(movements) {
		n = len(movements)
	}

	fmt.Println("Starting sketching...")
	drawer.Sketch(movements[:n])
}

func loadMovements() {
	fmt.Println("Loading Movements...")

	var err error
	path := "path_to_image.png"
	movements, err = imageprocessor.ProcessImage(path)
	if err != nil {
		fmt.Println("Error processing image: ", err)
	}
}

func NewApp() {
	application := app.New()
	window := application.NewWindow("Stencil")

	amountEntry := widget.NewEntry()
	amountEntry.SetText("0") // Default movement amount

	startBtn := widget.NewButton("Start Sketching", func() {
		startSketching(amountEntry.Text)
	})

	loadMovementsBtn := widget.NewButton("Load Movements", func() {
		loadMovements()
	})

	content := container.NewVBox(loadMovementsBtn, startBtn, amountEntry)

	window.SetContent(content)
	window.Resize(fyne.NewSize(500, 500))
	window.SetFixedSize(false)

	window.SetOnClosed(func() {
		application.Quit()
	})

	window.ShowAndRun()
}
