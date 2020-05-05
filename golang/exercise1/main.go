package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// CalcFuel calculates the mass of fuel needed to carry m.
// Negative fuel gets converted to 0.
func CalcFuel(mass int) (fuelMass int) {
	fuelMass = int(math.Floor(float64(mass)/3) - 2)
	if fuelMass < 0 {
		return 0
	}
	return
}

// CalcFuelRecHelper calculates the mass of fuel needed to carry m, adding the
// fueld needed to carry the extra fuel.
func CalcFuelRecHelper(mass int) int {
	if mass == 0 {
		return 0
	}
	return mass + CalcFuelRecHelper(CalcFuel(mass))
}

// CalcFuelRec calculates the mass of needed fuel and subtracts the initial item mass.
func CalcFuelRec(mass int) int {
	return CalcFuelRecHelper(mass) - mass
}

func main() {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)

	// Sum the fuel: floor(mass/3) - 2
	var sum int = 0
	for scanner.Scan() {
		sItemMass := scanner.Text()
		iItemMass, err := strconv.Atoi(sItemMass)
		if err != nil {
			panic(err)
		}
		sum += CalcFuel(iItemMass)
	}
	fmt.Printf("Simple fuel summary: %d\n", sum)

	// Reset reader.
	file.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(file)

	// Sum the fuel as above, recursively adding fuel weight until negative.
	var recSum int = 0
	for scanner.Scan() {
		sItemMass := scanner.Text()
		iItemMass, err := strconv.Atoi(sItemMass)
		if err != nil {
			panic(err)
		}
		recSum += CalcFuelRec(iItemMass)
	}
	fmt.Printf("Recursive fuel summary: %d\n", recSum)
}
