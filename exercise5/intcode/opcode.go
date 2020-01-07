package intcode

import "errors"

const (
	ADD    = 1
	MUL    = 2
	INPUT  = 3
	OUTPUT = 4
	HALT   = 99
)

type Opcode interface {
	Exec(c *Computer)
}

// OpBase is the base class for Opcode. Param[i] is a parameter of Reg[i].
type OpBase struct {
	Reg    []int
	Params []int
}

// NewOpcode reads a stream of Runes from a Computer and returns an executable Opcode.
func NewOpcode(c *Computer) (o Opcode, err error) {
	opID := c.next()
	opb := OpBase{}
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
		return &OpHalt{}, nil
	default:
		return nil, errors.New("unknown Opcode")
	}

}

type OpAdd struct {
	OpBase
}

func (op *OpAdd) Exec(c *Computer) {
	c.mem[op.Reg[2]] = c.mem[op.Reg[0]] + c.mem[op.Reg[1]]
}

type OpMul struct {
	OpBase
}

func (op *OpMul) Exec(c *Computer) {
	c.mem[op.Reg[2]] = c.mem[op.Reg[0]] * c.mem[op.Reg[1]]
}

type OpHalt struct{}

func (op *OpHalt) Exec(c *Computer) {
	c.halt = true
}

type OpInput struct{}

func (op *OpInput) Exec(c *Computer) {
	// Todo store input.
}

type OpOutput struct{}

func (op *OpOutput) Exec(c *Computer) {
	// Todo store output.
}
