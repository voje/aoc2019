package intcode_test

import (
	"testing"
	intcode "testcom.voje/aoc2019/exercise2/intcode"
)

var icc intcode.IntCodeComputer

func SetUp(t *testing.T) {
	icc = *intcode.NewIntCodeCompouter()
}

func TestSetReg(t *testing.T) {
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
}

func TestGetReg(t *testing.T) {
	SetUp(t)

}
