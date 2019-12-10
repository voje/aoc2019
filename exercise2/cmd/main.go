package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/voje/aoc2019/exercise2/intcode"
)

type Combo struct {
	Noun int `json:"n"`
	Verb int `json:"v"`
	Res  int `json:"r"`
}

func main() {
	// Read input data.
	file, err := os.Open("../data.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	icc := intcode.NewIntCodeComputer()
	slc, err := intcode.SliceFromReader(file)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	file.Close()

	icc.ReadSlice(slc)
	fmt.Printf("Input: %v\n", icc.GetRegs())
	icc.Execute(12, 2)
	fmt.Printf("Output: %v\n", icc.GetRegs())

	results := []Combo{}
	// Bruteforce the output.
	const goal int = 19690720
	var resCombo Combo
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			// Reset registers and execute.
			icc.ReadSlice(slc)
			icc.Execute(i, j)
			cmb := Combo{
				Noun: i,
				Verb: j,
				Res:  icc.GetReg(0),
			}
			if cmb.Res == goal {
				resCombo = cmb
			}
			results = append(results, cmb)
		}
	}

	for _, r := range results {
		fmt.Printf("%+v\n", r)
	}
	fmt.Printf("ResCombo: %+v\n", resCombo)

	// Save statistics for later usage (visualization?).
	// Create output file for some statistics on brute forcing.
	file, err = os.Create("../statistics.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	enc := json.NewEncoder(file)
	err = enc.Encode(&results)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
