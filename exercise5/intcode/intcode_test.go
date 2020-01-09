package intcode_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/voje/aoc2019/exercise5/intcode"
)

func TestAdd(t *testing.T) {
	// Should add up 2 and 3 and place result in [3] (last in array).
	c := intcode.NewComputer([]int{1, 4, 5, 0, 19, 23})
	t.Log(c.ToString())
	op, _ := intcode.NewOpcode(c)
	t.Log(op)
	op.Exec()
	t.Log(c.ToString())
	assert.Equal(t, 42, c.GetMem(0))
}

func TestMul(t *testing.T) {
	// Should add up 2 and 3 and place result in [3] (last in array).
	c := intcode.NewComputer([]int{2, 4, 5, 0, 3, 5})
	t.Log(c.ToString())
	op, _ := intcode.NewOpcode(c)
	t.Log(op)
	op.Exec()
	t.Log(c.ToString())
	assert.Equal(t, 15, c.GetMem(0))
}

func TestMul1(t *testing.T) {
	// If you precede an int with 0, it will use the octal numeric system. Avoid leading zeros.
	c := intcode.NewComputer([]int{102, 7, 3, 5, 99, 0})
	t.Log(c.ToString())
	op, _ := intcode.NewOpcode(c)
	t.Log(op)
	op.Exec()
	t.Log(c.ToString())
	assert.Equal(t, []int{102, 7, 3, 5, 99, 35}, c.DumpMem())
}

func TestOpcodeOperations(t *testing.T) {
	type Tsts struct {
		InMem  []int
		OutMem []int
	}
	tsts := []Tsts{
		{
			InMem:  []int{1, 0, 0, 0, 99},
			OutMem: []int{2, 0, 0, 0, 99},
		},
		{
			InMem:  []int{2, 3, 0, 3, 99},
			OutMem: []int{2, 3, 0, 6, 99},
		},
		{
			InMem:  []int{2, 4, 4, 5, 99, 0},
			OutMem: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			InMem:  []int{1102, 7, 3, 5, 99, 0},
			OutMem: []int{1102, 7, 3, 5, 99, 21},
		},
		{
			InMem:  []int{102, 7, 3, 5, 99, 0},
			OutMem: []int{102, 7, 3, 5, 99, 35},
		},
	}
	for _, tst := range tsts {
		c := intcode.NewComputer(tst.InMem)
		c.Run()
		assert.Equal(t, tst.OutMem, c.DumpMem())
	}
}

func TestInput(t *testing.T) {
	r := strings.NewReader("13")
	c := intcode.NewComputer([]int{3, 0, 99})
	c.Reader = r
	c.Run()
	t.Log(c.DumpMem())
}

func TestParseOpcodeID(t *testing.T) {
	type TestData struct {
		Opc    int
		Id     int
		Params []int
	}
	tests := []TestData{
		{
			Opc:    1002,
			Id:     2,
			Params: []int{0, 1},
		},
		{
			Opc:    02,
			Id:     2,
			Params: []int{},
		},
		{
			Opc:    2,
			Id:     2,
			Params: []int{},
		},
	}
	for _, tst := range tests {
		id, params := intcode.ParseOpCodeID(tst.Opc)
		assert.Equal(t, tst.Id, id)
		assert.Equal(t, tst.Params, params)
	}
}
