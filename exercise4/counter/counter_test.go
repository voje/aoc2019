package counter_test

import (
	"github.com/voje/aoc2019/exercise4/counter"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	c := counter.NewCounter(0)	
	for i := 0; i < 25; i++ {
		c.Tick()
		t.Log(c.ToString())
	}
	assert.Equal(t, c.ToString(), "25")
}

func TestDoubles(t *testing.T) {
	assert.Equal(t, 4, counter.NewCounter(11211132033).CheckAll())
}