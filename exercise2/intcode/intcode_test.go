package intcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/voje/aoc2019/exercise2/intcode"
)

var icc intcode.IntCodeComputer

func SetUp(t *testing.T) {
	icc = *intcode.NewIntCodeComputer()
}

func TestSetGetReg(t *testing.T) {
	SetUp(t)
	type pair struct {
		idx int
		val int
	}
	pairs := []pair{
		pair{0, 1},
		pair{1, 1},
		pair{10000, 1},
	}
	for _, p := range pairs {
		icc.SetReg(p.idx, p.val)
	}

	for _, p := range pairs {
		assert.Equal(t, p.val, icc.GetReg(p.idx))
	}
}

type testPair struct {
	In  []int
	Out []int
}

var testPairs = []testPair{
	{
		In:  []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		Out: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
	},
	{
		In:  []int{1, 0, 0, 0, 99},
		Out: []int{2, 0, 0, 0, 99},
	},
	{
		In:  []int{2, 3, 0, 3, 99},
		Out: []int{2, 3, 0, 6, 99},
	},
	{
		In:  []int{2, 4, 4, 5, 99, 0},
		Out: []int{2, 4, 4, 5, 99, 9801},
	},
	{
		In:  []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
		Out: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
	},
}

func TestExecutions(t *testing.T) {
	for _, pair := range testPairs {
		SetUp(t)
		icc.ReadSlice(pair.In)
		icc.Execute(pair.In[1], pair.In[2])
		assert.Equal(t, pair.Out, icc.GetRegs())
	}
}
