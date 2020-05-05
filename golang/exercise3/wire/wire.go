package wire

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/voje/aoc2019/exercise3/vector"
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
	wire.Steps = []*Step{}
	return wire
}

// NewWireFromString reads a string of format:
// R75,D30,R83,U83,L12,D49,R71,U7,L72
// and returns a *Wire.
func NewWireFromString(stxt string) (w *Wire) {
	w = NewWire()

	// Add initial step (does not matter, the grid is relative).
	w.AddStep(0, 0)

	// Loop all steps.
	stepTxts := strings.Split(stxt, ",")
	for _, stepTxt := range stepTxts {
		prevStep := w.Steps[len(w.Steps)-1]
		val, err := strconv.ParseFloat(stepTxt[1:], 64)
		if err != nil {
			fmt.Println(err)
		}
		switch stepTxt[0] {
		case 'U':
			w.AddStep(prevStep.X, prevStep.Y+val)
		case 'D':
			w.AddStep(prevStep.X, prevStep.Y-val)
		case 'L':
			w.AddStep(prevStep.X-val, prevStep.Y)
		case 'R':
			w.AddStep(prevStep.X+val, prevStep.Y)
		}
	}
	return w
}

// FindIntersections finds a list of intersections between two wires and returns them as pointers to (newly generated) Step structs.
func FindIntersections(w1, w2 *Wire) (intersections []*Step) {

	for i := 0; i < len(w1.Steps)-1; i++ {
		for j := 0; j < len(w2.Steps)-1; j++ {
			a := vector.Vector{X: w1.Steps[i].X, Y: w1.Steps[i].Y}
			b := vector.Vector{X: w1.Steps[i+1].X, Y: w1.Steps[i+1].Y}
			c := vector.Vector{X: w2.Steps[j].X, Y: w2.Steps[j].Y}
			d := vector.Vector{X: w2.Steps[j+1].X, Y: w2.Steps[j+1].Y}
			its := vector.Intersect(a, b, c, d)
			if its != nil {
				intersections = append(intersections, &Step{X: its.X, Y: its.Y})
			}
		}
	}
	return
}

func (w *Wire) AddStep(x, y float64) {
	var pathDist float64
	if len(w.Steps) == 0 {
		pathDist = 0
	} else {
		prevStep := w.Steps[len(w.Steps)-1]
		pathDist = prevStep.PathDist + math.Abs(prevStep.X-x) + math.Abs(prevStep.Y-y)
	}
	newStep := &Step{
		Wire:     w,
		X:        x,
		Y:        y,
		PathDist: pathDist,
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
	Wire     *Wire // Pointer to parent wire.
	X        float64
	Y        float64
	PathDist float64 // Accumulated disatnce of the traveled path along the wire.
}

func Manhattan(x1, y1, x2, y2 float64) float64 {
	return math.Abs(x2-x1) + math.Abs(y2-y1)
}

func (s *Step) CoordString() string {
	return fmt.Sprintf("%f-%f", s.X, s.Y)
}

func StepFromCoordString(s string) *Step {
	arr := strings.Split(s, "-")
	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[0])
	return &Step{
		X: float64(x),
		Y: float64(y),
	}
}

// Might be obsolete...
type Intersection struct {
	S1 *Step
	S2 *Step
}

// Follow will follow the wire until it reaches the desired point.
func (w *Wire) Follow(x, y float64) (dist float64, ok bool) {
	dist = 0
	for i := 0; i < len(w.Steps)-1; i++ {
		// Check if x, y lies between the two steps.
		stp1 := w.Steps[i]
		stp2 := w.Steps[i+1]
		ptDst, ok := PointBetween(stp1.X, stp1.Y, x, y, stp2.X, stp2.Y)
		if ok {
			return dist + ptDst, true
		}
		// Keep searching.
		dist += Manhattan(stp1.X, stp1.Y, stp2.X, stp2.Y)
	}
	return 0, false
}

// PointBetween checks if a point X2 lies between two edge points and (X1, X3) and returns the abs. distance from X1 to X2.
// Only works for this exercise (90 degree angles).
func PointBetween(x1, y1, x2, y2, x3, y3 float64) (dist float64, ok bool) {
	if x1 == x2 && x2 == x3 {
		if (y1 <= y2 && y2 <= y3) || (y3 <= y2 && y2 <= y1) {
			return math.Abs(y2 - y1), true
		}
	} else if y1 == y2 && y2 == y3 {
		if (x1 <= x2 && x2 <= x3) || (x3 <= x2 && x2 <= x1) {
			return math.Abs(x2 - x1), true
		}
	}
	return 0, false
}

// ShortestPathIntersection finds the shortest path between wire beginnings and an intersection.
// Returns a nil pointer if no intersection is found.
func ShortestPathIntersection(w1, w2 *Wire) (dist float64, point *Step) {
	its := FindIntersections(w1, w2)

	// The wires start at the same point -- that one should be omitted.
	if its[0].X == 0 && its[0].Y == 0 {
		its = its[1:]
	}

	dist = math.MaxFloat64
	point = nil
	for _, it := range its {
		w1PathDist, _ := w1.Follow(it.X, it.Y)
		w2PathDist, _ := w2.Follow(it.X, it.Y)
		dst := w1PathDist + w2PathDist
		if dst < dist {
			dist = dst
			point = &Step{X: it.X, Y: it.Y}
		}
	}
	return
}
