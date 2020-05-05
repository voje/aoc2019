package vector_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/voje/aoc2019/exercise3/vector"
)

type LinePair struct {
	A         vector.Vector
	B         vector.Vector
	C         vector.Vector
	D         vector.Vector
	Intersect *vector.Vector
}

var linePairs = []LinePair{
	// Simple X formation.
	LinePair{
		A:         vector.Vector{X: 1, Y: 1, Z: 0},
		B:         vector.Vector{X: 3, Y: 3, Z: 0},
		C:         vector.Vector{X: 1, Y: 3, Z: 0},
		D:         vector.Vector{X: 3, Y: 1, Z: 0},
		Intersect: &vector.Vector{X: 2, Y: 2, Z: 0},
	},
	LinePair{
		A:         vector.Vector{X: 1, Y: 1, Z: 0},
		B:         vector.Vector{X: 3, Y: 3, Z: 0},
		C:         vector.Vector{X: 4, Y: 1, Z: 0},
		D:         vector.Vector{X: 4, Y: 3, Z: 0},
		Intersect: nil,
	},
	// First line ends in second line.
	LinePair{
		A:         vector.Vector{X: 2, Y: 2, Z: 0},
		B:         vector.Vector{X: 4, Y: 2, Z: 0},
		C:         vector.Vector{X: 4, Y: 1, Z: 0},
		D:         vector.Vector{X: 4, Y: 3, Z: 0},
		Intersect: &vector.Vector{X: 4, Y: 2, Z: 0},
	},
	// Same lines. -- Apparently two same lines don't intersect?
	LinePair{
		A:         vector.Vector{X: 1, Y: 1, Z: 0},
		B:         vector.Vector{X: 3, Y: 3, Z: 0},
		C:         vector.Vector{X: 1, Y: 1, Z: 0},
		D:         vector.Vector{X: 3, Y: 3, Z: 0},
		Intersect: nil,
	},
}

func TestLinesIntersect(t *testing.T) {
	for _, lp := range linePairs {
		t.Log(lp)
		res := vector.Intersect(lp.A, lp.B, lp.C, lp.D)
		if lp.Intersect == nil {
			assert.Nil(t, res)
		} else {
			assert.Equal(t, lp.Intersect, vector.Intersect(lp.A, lp.B, lp.C, lp.D))
		}
	}
}
