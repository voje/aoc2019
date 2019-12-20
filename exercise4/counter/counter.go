package counter

import (
	"strconv"
)

type Counter struct {
	digits []int
}

func NewCounter(init int) *Counter {
	var digits []int
	if init == 0 {
		digits = []int{0}
	} else {
		for init != 0 {
			last := init % 10
			digits = append([]int{last}, digits...)
			init = init / 10
		}
	}
	return &Counter{
		digits: digits,
	}
}

func (c *Counter) ToString() (s string) {
	s = ""
	for _, d := range c.digits {
		s += strconv.Itoa(d)
	}
	return
}

func (c *Counter) Tick() int {
	return c.TickRec(len(c.digits) - 1)
}

// TickRec recursively ticks counter at a certain position.
// Returns the number of decimal pairs created in the process.
func (c *Counter) TickRec(pos int) (doubles int) {
	// Add digit if needed.
	if pos < 0 {
		c.digits[0] = 0
		c.digits = append([]int{1}, c.digits...)
		return
	}

	c.digits[pos]++
	for i := pos + 1; i < len(c.digits); i++ {
		c.digits[i] = c.digits[pos]
	}

	// Change digit if we reach 10.
	if c.digits[pos] == 10 {
		c.digits[pos] = 0
		c.TickRec(pos - 1)
	}
	return
}

// HasDoubles will return true if the digits contain pairs.
func (c *Counter) HasDoubles() bool {
	for i := 0; i < len(c.digits)-1; i++ {
		if c.digits[i] == c.digits[i+1] {
			return true
		}
	}
	return false
}
