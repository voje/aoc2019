package main

import(
	"fmt"
	"strconv"
)

func main() {
	// Completely overenginnered my solution... going with a simple main.go now...
	start := 206938
	end := 679128
	count := 0
	for num := start; num <= end; num++ {
		if check(num) {
			count ++
		}
	}
	fmt.Printf("Found %d numbers that match the criteria.\n", count)
}

// In one loop, check that there is at leas one same-digit pair and that all digits are ascending.  
func check(n int) bool {
	hasPair := false
	sn := strconv.Itoa(n)
	for i := 0; i < len(sn) - 1; i++ {
		// Check that all digits are equal or greater than their left neighbour.  
		if sn[i + 1] > sn[i] {
			return false
		}
		if sn[i] == sn [i + 1] {
			hasPair = true
		}
	}
	return hasPair
}