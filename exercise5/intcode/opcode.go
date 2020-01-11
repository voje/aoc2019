package intcode

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Opcode IDs
const (
	ADD    = 1
	MUL    = 2
	INPUT  = 3
	OUTPUT = 4
	JUMPIFTRUE = 5
	JUMPIFFALSE = 6
	LT = 7
	EQ = 8
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
		opb.Reg = append(opb.Reg, c.next(), c.next(), c.next())
		return &OpAdd{
			OpBase: opb,
		}, nil
	case MUL:
		opb.Reg = append(opb.Reg, c.next(), c.next(), c.next())
		return &OpMul{
			OpBase: opb,
		}, nil
	case HALT:
		return &OpHalt{
			OpBase: opb,
		}, nil
	case INPUT:
		opb.Reg = append(opb.Reg, c.next())
		return &OpInput{
			OpBase: opb,
		}, nil
	case OUTPUT:
		opb.Reg = append(opb.Reg, c.next())
		return &OpOutput{
			OpBase: opb,
		}, nil
	case JUMPIFTRUE:
		opb.Reg = append(opb.Reg, c.next(), c.next())
		return &OpJumpIfTrue{
			OpBase: opb,
		}, nil
	case JUMPIFFALSE:
		opb.Reg = append(opb.Reg, c.next(), c.next())
		return &OpJumpIfFalse{
			OpBase: opb,
		}, nil
	case LT:
		opb.Reg = append(opb.Reg, c.next(), c.next(), c.next())
		return &OpLessThan{
			OpBase: opb,
		}, nil
	case EQ:
		opb.Reg = append(opb.Reg, c.next(), c.next(), c.next())
		return &OpEquals{
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

type OpInput struct {
	OpBase
}

func (op *OpInput) Exec() {
	reader := bufio.NewReader(op.c.Reader)
	strInput, _ := reader.ReadString('\n')
	strInput = strings.Trim(strInput, "\n")
	intInput, err := strconv.Atoi(strInput)
	if err != nil {
		panic(err)
	}
	op.c.mem[op.Reg[0]] = intInput
}

type OpOutput struct {
	OpBase
}

func (op *OpOutput) Exec() {
	fmt.Fprintln(op.c.Writer, op.c.mem[op.Reg[0]])
}

type OpJumpIfTrue struct {
	OpBase
}

func (op *OpJumpIfTrue) Exec() {
	if op.Reg[0] != 0 {
		op.c.pc = op.Reg[1]
	}
}

type OpJumpIfFalse struct {
	OpBase
}

func (op *OpJumpIfFalse) Exec() {
	if op.Reg[0] == 0 {
		op.c.pc = op.Reg[1]
	}
}

type OpLessThan struct {
	OpBase
}

func (op *OpLessThan) Exec() {
	if op.Reg[0] < op.Reg[1]
}

type OpEquals struct {
	OpBase
}

func (op *OpEquals) Exec() {
	// TODO
}
