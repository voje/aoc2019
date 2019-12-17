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
	Steps      []*Step
	StepLookup map[string]*Step // Points to steps in Steps.
}

func NewWire() *Wire {
	wire := &Wire{
		StepLookup: make(map[string]*Step),
	}
	steps := []*Step{}
	initStep := &Step{Wire: wire, X: 0, Y: 0, PathDist: 0}
	steps = append(steps, initStep)
	wire.Steps = steps
	wire.StepLookup[initStep.CoordString()] = initStep
	return wire
}

func (w *Wire) AddStep(x, y float64) {
	var prevStep *Step
	prevStep = w.Steps[len(w.Steps)-1]
	newStep := &Step{
		Wire:	  w,
		X:        x,
		Y:        y,
		PathDist: prevStep.PathDist + math.Abs(prevStep.X-x) + math.Abs(prevStep.Y-y),
	}

	w.Steps = append(w.Steps, newStep)
	// Only add step if the map field is empty -- if the line intersects itself, we want a pointer to the
	// step closer to init.  
	if _, ok := w.StepLookup[newStep.CoordString()]; !ok {
		w.StepLookup[newStep.CoordString()] = newStep
	}
}

// Step holds 2D coordinates, distance from origin and intersections.
type Step struct {
	Wire 		  *Wire  // Pointer to parent wire. 
	X             float64
	Y             float64
	PathDist      float64  // Accumulated disatnce of the traveled path along the wire.  
}

func (s *Step) CoordString () string {
	return fmt.Sprintf("%f-%f", s.X, s.Y)
}

type Intersection struct {
	S1 *Step
	S2 *Step
}

// Follow will follow the wire until it reaches the desired point. 
func (w *Wire) Folow(x, y float64) (dist float64, ok bool) {
	var dist float64 = 0
	for i := 0; i < len(w.Steps) - 1; i++ {
		// Check if x, y lies between the two steps. 
		stp1 := w.Steps[i]
		stp2 := w.Steps[i + 1]
		if x == stp1.X && x == stp2.X {
			if y >= stp1.Y && y <=stp2.Y {
				return stp2.Y - stp1.Y, true
			}
		}
	}
}
