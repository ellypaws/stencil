package drawer

import (
	"github.com/ellypaws/stencil/image-processor"
	"github.com/go-vgo/robotgo"
	"time"
)

// A simple delay to simulate time taken to move between points and give the drawing a more "human" touch
const movementDelay = 10 * time.Millisecond

var Speed float64 = 5.0

func Sketch(movements []*imageprocessor.Movement) {
	robotgo.MouseSleep = 10

	for _, move := range movements {
		switch move.Type {
		case "move":
			robotgo.MoveSmooth(move.Point.X, move.Point.Y, 1.0, Speed)
			robotgo.Toggle("up", "left")
		case "draw":
			robotgo.MoveSmooth(move.Point.X, move.Point.Y, 1.0, Speed)
			robotgo.Toggle("down", "left")
		}
		time.Sleep(movementDelay)
	}
	robotgo.Toggle("up", "left")
}
