package main

import (
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
			fmt.Println(num)
			count++
		}
	}
	fmt.Printf("Found %d numbers that match the criteria.\n", count)
}

// In one loop, check that there is at leas one same-digit pair and that all digits are ascending.
func check(n int) bool {
	hasPair := false
	sn := strconv.Itoa(n)
	for i := 0; i < len(sn)-1; i++ {
		// Check that all digits are equal or greater than their left neighbour.
		if sn[i+1] < sn[i] {
			return false
		}

		// Check if there's a size 2 digit group.
		grpSize := digitGroupSize(sn, i)
		/*
			if grpSize > 1 {
				fmt.Printf("n: %s, size: %d\n", sn, grpSize)
			}
		*/
		if grpSize == 2 {
			hasPair = true
		}
	}
	return hasPair
}

func digitGroupSize(num string, idx int) int {
	comp := num[idx]
	grp := 1 // Starting with idx twice, prevent double count.
	for i := idx + 1; i < len(num); i++ {
		if comp == num[i] {
			grp++
		}
	}
	for i := idx - 1; i >= 0; i-- {
		if comp == num[i] {
			grp++
		}
	}
	return grp
}
