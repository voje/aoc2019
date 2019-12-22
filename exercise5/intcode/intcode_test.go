package intcode_test

import (
	"github.com/voje/aoc2019/exercise5/intcode"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Should add up 2 and 3 and place result in [3] (last in array).  
	c := intcode.NewComputer([]int{1,4,5,0,19,23})
	t.Log(c.ToString())
	op := c.ReadOpcode()
	t.Log(op)
	op.Exec(c)
	t.Log(c.ToString())
	assert.Equal(t, 42, c.GetMem(0))
}

func TestMul(t *testing.T) {
	// Should add up 2 and 3 and place result in [3] (last in array).  
	c := intcode.NewComputer([]int{2,4,5,0,3,5})
	t.Log(c.ToString())
	op := c.ReadOpcode()
	t.Log(op)
	op.Exec(c)
	t.Log(c.ToString())
	assert.Equal(t, 15, c.GetMem(0))
}

func TestRun1(t *testing.T) {
	c := intcode.NewComputer([]int{1,0,0,0,99})
	c.Run()
	t.Log(c.ToString())
	assert.Equal(t, 2, c.GetMem(0))
}

func TestRun2(t *testing.T) {
	c := intcode.NewComputer([]int{2,3,0,3,99})
	c.Run()
	t.Log(c.ToString())
	assert.Equal(t, 6, c.GetMem(3))
}

func TestRun3(t *testing.T) {
	c := intcode.NewComputer([]int{2,4,4,5,99,0})
	c.Run()
	t.Log(c.ToString())
	assert.Equal(t, 9801, c.GetMem(5))
}

func TestRun4(t *testing.T) {
	c := intcode.NewComputer([]int{1,1,1,4,99,5,6,0,99})
	c.Run()
	t.Log(c.ToString())
	assert.Equal(t, 30, c.GetMem(0))
}
