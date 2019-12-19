package counter

import "fmt"

type Counter struct {
	digits []int
}

func NewCounter(init int) *Counter {
	digits := []int{}
	for init != 0 {
		last := init % 10
		digits = append([]int{last}, digits...)
		init = init / 10
	}
	return &Counter{
		digits: digits,
	}
}

func (c *Counter) ToString() string {
	return fmt.Sprintf("%+v", c.digits)
}

func (c *Counter) Tick() {
	c.TickRec(len(c.digits) - 1)
}

// Recursively ticks counter at a certain position.
func (c *Counter) TickRec(pos int) {
	c.digits[pos]++

	// Add digit if needed.
	if pos < 0 {
		c.digits[0] = 0
		c.digits = append([]int{1}, c.digits...)
		return
	}

	if c.digits[pos] == 10 {
		c.digits[pos] = 0
		c.TickRec(pos - 1)
		return
	}
}
