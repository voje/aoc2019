package wire

import (
	"fmt"
	"math"
)

var WireCount int = 0

// Wire holds information about a wire being run along a coordinate system.
// It consists of Steps (coordinates).
// Steps are stored in an array. Pointers to those objects are stored in a map for quick lookup.
type Wire struct {
	Name       string
	Steps      []*Step
	StepLookup map[string]*Step // Points to steps in Steps.
}

func (w *Wire) NewWire() *Wire {
	return &Wire{
		Steps:      []*Step{},
		StepLookup: make(map[string]*Step),
	}
}

func (w *Wire) AddStep(x, y float64) {
	var prevStep *Step
	if len(w.Steps) == 0 {
		prevStep = &Step{
			X: 0,
			Y: 0,
		}
	} else {
		prevStep = w.Steps[len(w.Steps)-1]
	}
	newStep := &Step{
		X:        x,
		Y:        y,
		PathDist: prevStep.PathDist + math.Abs(prevStep.X-x) + math.Abs(prevStep.Y-y),
	}

	w.Steps = append(w.Steps, newStep)
	w.StepLookup[fmt.Sprintf("%f-%f", x, y)] = newStep
}

// Step holds 2D coordinates, distance from origin and intersections.
type Step struct {
	X             float64
	Y             float64
	PathDist      float64
	Intersections bool // todo: find a smart way to store these
}
