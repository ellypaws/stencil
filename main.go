package main

import (
	"fmt"
	"github.com/ellypaws/stencil/drawer"
	imageprocessor "github.com/ellypaws/stencil/image-processor"
)

func main() {
	path := "path_to_image.png"
	movements, err := imageprocessor.ProcessImage(path)
	if err != nil {
		fmt.Println("Error processing image:", err)
		return
	}

	// Now, you'd use these movements in your drawer package.
	drawer.Sketch(movements)
}
