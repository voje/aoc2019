package main

import (
	"fmt"
	"io"
	"os"
)

// FuelCounterUpper counts up the needed fuel.
type FuelCounterUpper struct {
	r io.Reader
}

// LineReader reads an input streams and outputs line-per-line.
// (Yes, I am aware of the bufio package.)
type LineReader struct {
	r io.Reader
}

// Read has a caveat. If line length exceeds buffer length, the program will have undefined behaviour.
func (lr *LineReader) Read(b []byte) (n int, err error) {
	lineBuffer := buf
	n, err := lr.r.Read(buf)
}

func main() {
	file, _ := os.Open("data.txt")
	fcu := FuelCounterUpper{r: file}
	fmt.Println(fcu.CalcFuel())
}
