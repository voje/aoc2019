package wire_test

import (
	"testing"

	"github.com/voje/aoc2019/exercise3/wire"
	"gotest.tools/assert"
)

func SetUp(t *testing.T) (w *wire.Wire) {
	w = wire.NewWire()

	w.AddStep(2, 1)
	w.AddStep(2, 3)
	w.AddStep(-1, 3)
	w.AddStep(-2, 3)
	w.AddStep(-2, 1)
	w.AddStep(-3, 1)
	w.AddStep(-3, 2)
	w.AddStep(-1, 2)
	w.AddStep(-1, -2)
	w.AddStep(-4, -2)
	w.AddStep(-4, 2)

	return w
}

func TestFollow(t *testing.T) {
	w := SetUp(t)

	// Ceck for negative (points outside of wire).
	_, ok := w.Follow(100, 100)
	assert.Equal(t, false, ok)

	// Check for positive.
	var dist float64
	dist, _ = w.Follow(2, 2)
	assert.Equal(t, dist, 1.0)

	dist, _ = w.Follow(1, 3)
	assert.Equal(t, dist, 3.0)

	dist, _ = w.Follow(-2, 2)
	assert.Equal(t, dist, 7.0)

	dist, _ = w.Follow(-2, -2)
	assert.Equal(t, dist, 17.0)

	dist, _ = w.Follow(-4, 0)
	assert.Equal(t, dist, 21.0)
}

func TestFollow1(t *testing.T) {
	sw1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	w1 := wire.NewWireFromString(sw1)

	sw2 := "U62,R66,U55,R34,D71,R55,D58,R83"
	w2 := wire.NewWireFromString(sw2)

	dist, _ := wire.ShortestPathIntersection(w1, w2)
	assert.Equal(t, dist, 610.0)
}

func TestFollow2(t *testing.T) {
	sw1 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	w1 := wire.NewWireFromString(sw1)

	sw2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
	w2 := wire.NewWireFromString(sw2)

	dist, _ := wire.ShortestPathIntersection(w1, w2)
	assert.Equal(t, dist, 410.0)
}
