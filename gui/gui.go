package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ellypaws/stencil/drawer"
	imageprocessor "github.com/ellypaws/stencil/image-processor"
)

var movements []*imageprocessor.Movement
var slider *widget.Slider
var sliderSpeed *widget.Slider
var labelMovementsInfo *widget.Label

func startSketching() {
	if len(movements) == 0 {
		fmt.Println("Movements not loaded yet.")
		return
	}

	movementsAmount := int(slider.Value)
	if movementsAmount > len(movements) {
		movementsAmount = len(movements)
	}

	drawer.Speed = sliderSpeed.Value

	fmt.Println("Starting sketching...")
	drawer.Sketch(movements[:movementsAmount])
}

func loadMovements() {
	fmt.Println("Loading Movements...")

	var err error
	path := "path_to_image.png"
	movements, err = imageprocessor.ProcessImage(path)
	if err != nil {
		fmt.Println("Error processing image: ", err)
		return
	}

	slider = widget.NewSlider(0, float64(len(movements)))
	slider.Value = float64(len(movements))

	labelMovementsInfo.SetText(fmt.Sprintf("Loaded %d movements", len(movements)))
}

func NewApp() {
	application := app.New()
	window := application.NewWindow("Stencil")
	sliderSpeed = widget.NewSlider(0.1, 10.0)
	sliderSpeed.Value = 0.5
	labelMovementsInfo = widget.NewLabel("")

	loadMovements()

	content := container.NewVBox(
		widget.NewButton("Load Movements", loadMovements),
		widget.NewLabel("Number of movements:"),
		slider,
		labelMovementsInfo,
		widget.NewLabel("Speed:"),
		sliderSpeed,
		widget.NewButton("Start Sketching", startSketching),
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(500, 500))
	window.SetFixedSize(false)

	window.SetOnClosed(func() {
		application.Quit()
	})

	window.ShowAndRun()
}
