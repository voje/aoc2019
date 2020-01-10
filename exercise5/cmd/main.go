package main

import (
	"os"

	"github.com/voje/aoc2019/exercise5/intcode"
)

func main() {
	f, err := os.Open("../data.txt")
	if err != nil {
		panic(err)
	}
	c, _ := intcode.NewComputerFromReader(f)
	// fmt.Println(c.DumpMem())
	c.Run()
}
