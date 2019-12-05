package main

import (
	"fmt"
	"io"
	"os"
)

// FuelCounterUpper counts up the needed fuel.
type FuelCounterUpper struct {
	r         io.Reader
	FuelCount int
}

func main() {
	file, _ := os.Open("data.txt")
	fcu := FuelCounterUpper{r: file}
	fmt.Println(fcu.FuelCount)
}
