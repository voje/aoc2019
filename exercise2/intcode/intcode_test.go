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

func TestExecutions(t *testing.T) {
	SetUp(t)
	in1 := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	out1 := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
	icc.ReadSlice(in1)
	t.Logf("InputCod: %+v", icc.GetRegs())
	icc.Execute()
	t.Logf("ExecdRes: %+v", icc.GetRegs())
	t.Logf("ValidRes: %+v", out1)
	assert.Equal(t, out1, icc.GetRegs())
}
