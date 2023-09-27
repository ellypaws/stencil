package drawer

import (
	"github.com/ellypaws/stencil/image-processor"
	"github.com/go-vgo/robotgo"
	"time"
)

// A simple delay to simulate time taken to move between points and give the drawing a more "human" touch
const movementDelay = 10 * time.Millisecond

// Sketch performs drawing based on a series of movements using robotgo.
// Note: In the latest versions of robotgo, methods like MoveMouse, MoveMouseSmooth, etc. are deprecated.
// Basically remove the word "Mouse" from the method name and use it. e.g., MoveMouse -> Move
// It is recommended to use the Move() method for moving the mouse.
func Sketch(movements []imageprocessor.Movement) {
	robotgo.MouseSleep = 10 // Introduce some delay between mouse operations for smoother movement

	for _, move := range movements {
		switch move.Type {
		case "move":
			robotgo.MoveSmooth(move.Point.X, move.Point.Y, 1.0, 5.0) // Using MoveSmooth for more "human-like" movement
			robotgo.Toggle("up", "left")                             // Ensure mouse button is released
		case "draw":
			robotgo.MoveSmooth(move.Point.X, move.Point.Y, 1.0, 5.0)
			robotgo.Toggle("down", "left") // Press mouse button to draw
		}
		time.Sleep(movementDelay)
	}
	robotgo.Toggle("up", "left") // Ensure mouse button is released at the end
}
