package counter_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/voje/aoc2019/exercise4/counter"
	"testing"
)

func TestCounter(t *testing.T) {
	c := counter.NewCounter(0)
	for i := 0; i < 25; i++ {
		c.Tick()
		t.Log(c.ToString())
	}
	assert.Equal(t, c.ToString(), "25")
}
