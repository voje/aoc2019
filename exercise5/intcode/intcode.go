package intcode

import (
	"fmt"
)

const (
	ADD = 1
	MUL = 2
	INPUT = 3
	OUTPUT = 4
	HALT = 99
)

type Opcode interface {
	Exec(c *Computer)
}

type OpAdd struct {
	R1 int
	R2 int
	R3 int
}

func (op *OpAdd) Exec(c *Computer) {
	c.mem[op.R3] = c.mem[op.R1] + c.mem[op.R2]
}

type OpMul struct {
	R1 int
	R2 int
	R3 int
}

func (op *OpMul) Exec(c *Computer) {
	c.mem[op.R3] = c.mem[op.R1] * c.mem[op.R2]
}

type OpHalt struct {}

func (op *OpHalt) Exec(c *Computer) {
	c.halt = true
}

type OpInput struct {}

func (op *OpInput) Exec(c *Computer) {
	// Todo store input.
}

type OpOutput struct {}

func (op *OpOutput) Exec(c *Computer) {
	// Todo store output.
}

type Computer struct {
	mem []int
	pc int
	halt bool
}

func NewComputer(mem []int) *Computer {
	return &Computer {
		mem: mem,
		pc: 0,
		halt: false,
	}
}

func (c *Computer) ToString() string {
	return fmt.Sprintf("%+v\n", *c)
}

func (c *Computer) next() int {
	c.pc += 1
	return c.mem[c.pc - 1]
}

func (c* Computer) GetMem(i int) int {
	return c.mem[i]
}

func (c* Computer) Run() {
	for !c.halt {
		op := c.ReadOpcode()
		op.Exec(c)
	}
}

// ReadOpcode reads a sequence of ints and returns an Opcode.
// Returns an error if the sequence is nota valid opcode. 
func (c *Computer) ReadOpcode() Opcode {
	opId := c.next()
	switch opId {
	case ADD:
		return &OpAdd{
			R1: c.next(),
			R2: c.next(),
			R3: c.next(),
		}
	case MUL:
		return &OpMul{
			R1: c.next(),
			R2: c.next(),
			R3: c.next(),
		}
	case HALT:
		return &OpHalt{}
	default:
		return nil
	}
}
