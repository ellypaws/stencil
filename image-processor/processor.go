package imageprocessor

import (
	"github.com/nfnt/resize"
	"image"
	"image/color"
	_ "image/png"
	"os"
)

// Movement holds the type and coordinates of a movement.
type Movement struct {
	Type  string      // "move" or "draw"
	Point image.Point // the target point
}

// ProcessImage processes the input image, identifies edges, and returns movements.
func ProcessImage(filePath string) ([]Movement, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// Resize image for performance (optional)
	small := resize.Resize(100, 0, img, resize.Bilinear)

	movements := make([]Movement, 0)
	bounds := small.Bounds()

	lastPixelWasBlack := false

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			col := color.GrayModel.Convert(small.At(x, y)).(color.Gray)
			if col.Y > 128 {
				if lastPixelWasBlack {
					movements = append(movements, Movement{Type: "move", Point: image.Point{X: x, Y: y}})
					lastPixelWasBlack = false
				}
			} else {
				if !lastPixelWasBlack {
					movements = append(movements, Movement{Type: "draw", Point: image.Point{X: x, Y: y}})
					lastPixelWasBlack = true
				}
			}
		}
	}

	return movements, nil
}
