package intcode

import (
	"fmt"
)

type Computer struct {
	mem  []int
	pc   int
	halt bool
}

func NewComputer(mem []int) *Computer {
	return &Computer{
		mem:  mem,
		pc:   0,
		halt: false,
	}
}

func (c *Computer) ToString() string {
	return fmt.Sprintf("%+v\n", *c)
}

func (c *Computer) next() int {
	c.pc += 1
	return c.mem[c.pc-1]
}

func (c *Computer) GetMem(i int) int {
	return c.mem[i]
}

func (c *Computer) DumpMem() []int {
	return c.mem
}

func (c *Computer) Run() error {
	for !c.halt {
		op, err := NewOpcode(c)
		if err != nil {
			return err
		}
		op.Exec()
	}
	return nil
}

func (c *Computer) Halt() {
	c.halt = true
}
