package intcode

import (
	"errors"
	"strconv"
)

// Opcode IDs
const (
	ADD    = 1
	MUL    = 2
	INPUT  = 3
	OUTPUT = 4
	HALT   = 99
)

// Parameter modes
const (
	POS       = 0
	IMMEDIATE = 1
)

type Opcode interface {
	Exec()
	GetAt(int) int
}

// OpBase is the base class for Opcode. Param[i] is a parameter of Reg[i].
type OpBase struct {
	c      *Computer
	Reg    []int
	Params []int
}

// Based on the parameter Params[i], get value from Reg[i].
func (o *OpBase) GetAt(i int) int {
	if i < len(o.Params) && o.Params[i] == IMMEDIATE {
		return o.Reg[i]
	}
	return o.c.GetMem(o.Reg[i])
}

// ParseOpCodeID parses an opcode ABCD.
// Last two digits (CD) prepresent the Opcode ID.
// Remaining digits from right to left represent parameters for opcode arguments (registers).
// B -> Reg[0], A -> Reg[1]
func ParseOpCodeID(opID int) (id int, params []int) {
	sOpID := strconv.Itoa(opID)

	// Return without params.
	if len(sOpID) < 2 {
		return opID, []int{}
	}

	// Return with params.
	partID := sOpID[(len(sOpID) - 2):]
	id, _ = strconv.Atoi(string(partID))
	partParam := sOpID[:(len(sOpID) - 2)]

	for i := len(partParam) - 1; i >= 0; i-- {
		// Convert ascii one-digit rune to int.
		r := int([]rune(partParam)[i]) - 48
		params = append(params, r)
	}
	return id, params
}

// NewOpcode reads a stream of Runes from a Computer and returns an executable Opcode.
func NewOpcode(c *Computer) (o Opcode, err error) {
	opID := c.next()
	opb := OpBase{
		c: c,
	}
	opID, opb.Params = ParseOpCodeID(opID)

	// Store parameters.
	switch opID {
	case ADD:
		for i := 0; i < 3; i++ {
			opb.Reg = append(opb.Reg, c.next())
		}
		return &OpAdd{
			OpBase: opb,
		}, nil
	case MUL:
		for i := 0; i < 3; i++ {
			opb.Reg = append(opb.Reg, c.next())
		}
		return &OpMul{
			OpBase: opb,
		}, nil
	case HALT:
		return &OpHalt{
			OpBase: opb,
		}, nil
	default:
		return nil, errors.New("unknown Opcode")
	}

}

type OpAdd struct {
	OpBase
}

func (op *OpAdd) Exec() {
	op.c.mem[op.Reg[2]] = op.GetAt(0) + op.GetAt(1)
}

type OpMul struct {
	OpBase
}

func (op *OpMul) Exec() {
	op.c.mem[op.Reg[2]] = op.GetAt(0) * op.GetAt(1)
}

type OpHalt struct {
	OpBase
}

func (op *OpHalt) Exec() {
	op.c.Halt()
}

type OpInput struct{}

func (op *OpInput) Exec(c *Computer) {
	// Todo store input.
}

type OpOutput struct{}

func (op *OpOutput) Exec(c *Computer) {
	// Todo store output.
}
