package intcode

type IntCodeComputer struct {
	regs []int
}

// NewIntCodeComputer reates a new instance of IntCodeComputer.
func NewIntCodeCompouter() *IntCodeComputer {
	return &IntCodeComputer{regs: make([]int, 10)}
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
