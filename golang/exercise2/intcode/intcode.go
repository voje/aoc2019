package intcode

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type IntCodeComputer struct {
	regs []int
	pc   int
}

const (
	ADD int = 1
	MUL int = 2
	END int = 99
)

type Opcode struct {
	Id     int
	Arg1   int
	Arg2   int
	ResIdx int
}

// EOP indicates the end of the program.
var EOP = errors.New("End of program.")

// NewIntCodeComputer reates a new instance of IntCodeComputer.
func NewIntCodeComputer() *IntCodeComputer {
	return &IntCodeComputer{regs: make([]int, 10), pc: 0}
}

// Partially borrowed from StackOverflow
func csvSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the next comma.
	if i := strings.Index(string(data), ","); i >= 0 {
		return i + 1, data[0:i], nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), data, nil
	}

	return
}

// ReadSlice resets all registers to 0 and reads in a new slice.
func (icc *IntCodeComputer) ReadSlice(s []int) {
	icc.pc = 0
	for i := range icc.regs {
		icc.SetReg(i, 0)
	}
	for i, el := range s {
		icc.SetReg(i, el)
	}
}

func TrimTrailingZeros(slc []int) (newSlc []int) {
	lastNonZero := 0
	for i, e := range slc {
		if e != 0 {
			lastNonZero = i
		}
	}
	newSlc = slc[:lastNonZero+1]
	return newSlc
}

func SliceFromReader(r io.Reader) ([]int, error) {
	s := bufio.NewScanner(r)
	s.Split(csvSplitFunc)
	slc := []int{}
	for s.Scan() {
		iText, err := strconv.Atoi(strings.Trim(s.Text(), "\n"))
		if err != nil {
			return nil, err
		}
		slc = append(slc, iText)
	}
	return slc, nil
}

func (icc *IntCodeComputer) SetReg(i, val int) {
	if i >= cap(icc.regs) {
		extension := make([]int, i)
		icc.regs = append(icc.regs, extension...)
	}
	icc.regs[i] = val
}

func (icc *IntCodeComputer) GetReg(i int) int {
	return icc.regs[i]
}

func (icc *IntCodeComputer) NextOpcode() (*Opcode, error) {
	if icc.GetReg(icc.pc) == END {
		return nil, EOP
	}
	// Increment program counter.
	icc.pc += 4
	return &Opcode{
		Id:     icc.GetReg(icc.pc - 4),
		Arg1:   icc.GetReg(icc.pc - 3),
		Arg2:   icc.GetReg(icc.pc - 2),
		ResIdx: icc.GetReg(icc.pc - 1),
	}, nil
}

// Execute executes the command, using noun and verb.
// Noun is the command at address 1, verb is the address at address 2.
func (icc *IntCodeComputer) Execute(noun, verb int) {
	icc.SetReg(1, noun)
	icc.SetReg(2, verb)
	for {
		opc, err := icc.NextOpcode()
		if err == EOP {
			return
		}
		icc.ExecOpcode(opc)
	}
}

// ExecOpcode performs the Opcode operation and stores the result to icc registers.
func (icc *IntCodeComputer) ExecOpcode(o *Opcode) {
	switch o.Id {
	case ADD:
		icc.SetReg(o.ResIdx, icc.GetReg(o.Arg1)+icc.GetReg(o.Arg2))
	case MUL:
		icc.SetReg(o.ResIdx, icc.GetReg(o.Arg1)*icc.GetReg(o.Arg2))
	}
}

func (icc *IntCodeComputer) GetRegs() []int {
	regs := []int{}
	for _, r := range icc.regs {
		regs = append(regs, r)
	}
	return TrimTrailingZeros(regs)
}
