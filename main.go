package main

import (
	"github.com/ellypaws/stencil/drawer"
	imageprocessor "github.com/ellypaws/stencil/image-processor"
)

func main() {
	path := "path_to_image.png" // Replace with your image path
	points := imageprocessor.ProcessImage(path)

	// Start drawing
	drawer.Sketch(points)

	// Erase (if needed, perhaps based on certain conditions)
	drawer.Erase(points)
}
