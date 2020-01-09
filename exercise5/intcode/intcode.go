package intcode

import (
	"fmt"
	"io"
	"os"
	"encoding/csv"
	"strconv"
)

type Computer struct {
	Reader io.Reader // Useful for simulating user input with io.Writer.
	mem    []int
	pc     int
	halt   bool
}

func NewComputer(mem []int) *Computer {
	return &Computer{
		Reader: os.Stdin,
		mem:  mem,
		pc:   0,
		halt: false,
	}
}

func NewComputerFromReader(r io.Reader) (c *Computer, err error) {
	var mem []int
	reader := csv.NewReader(r)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		for _ , strNum := range line {
			intNum, _ := strconv.Atoi(strNum)
			mem = append(mem, intNum)
		}
	}
	c = NewComputer(mem)
	return c, nil
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
