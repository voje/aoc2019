package intcode_test

import (
	"testing"
	"github.com/voje/aoc2019/exercise2/intcode"
        "github.com/stretchr/testify/assert"
)

var icc intcode.IntCodeComputer

func SetUp(t *testing.T) {
	icc = *intcode.NewIntCodeCompouter()
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

        for _, p:= range pairs {
            assert.Equal(t, p.val, icc.GetReg(p.idx))
        }
}
