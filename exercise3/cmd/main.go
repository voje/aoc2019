package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/voje/aoc2019/exercise3/vector"
)

func main() {
	file, err := os.Open("../data.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Start two paths, both begin at zero.
	var paths [][]vector.Vector

	// Read two lines of file and fill the paths.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newPath := []vector.Vector{vector.Vector{X: 0, Y: 0, Z: 0}}
		pathCoords := strings.Split(scanner.Text(), ",")
		for _, coord := range pathCoords {
			// i should point to the last-added coordinate in the list.
			prev := newPath[len(newPath)-1]
			val, err := strconv.ParseFloat(coord[1:], 64)
			if err != nil {
				fmt.Println(err)
			}
			switch coord[0] {
			case 'U':
				newPath = append(newPath, vector.Vector{X: prev.X, Y: prev.Y + val, Z: prev.Z})
			case 'D':
				newPath = append(newPath, vector.Vector{X: prev.X, Y: prev.Y - val, Z: prev.Z})
			case 'L':
				newPath = append(newPath, vector.Vector{X: prev.X - val, Y: prev.Y, Z: prev.Z})
			case 'R':
				newPath = append(newPath, vector.Vector{X: prev.X + val, Y: prev.Y, Z: prev.Z})
			}
		}
		paths = append(paths, newPath)
	}

	fmt.Printf("paths:\n%+v\n", paths)

	// Find all intersections.
	var intersections []vector.Vector
	for i := 0; i < len(paths[0])-1; i++ {
		for j := 0; j < len(paths[1])-1; j++ {
			intrsct := vector.Intersect(paths[0][i], paths[0][i+1], paths[1][j], paths[1][j+1])
			if intrsct != nil {
				intersections = append(intersections, *intrsct)
			}
		}
	}

	fmt.Printf("intersections:\n%+v\n", intersections)

	// Find the intersection closest to (0, 0).
	closest := vector.Vector{X: math.MaxFloat64, Y: math.MaxFloat64, Z: 0}
	zero := vector.Vector{X: 0, Y: 0, Z: 0}
	for _, p := range intersections {
		if vector.Manhattan(p, zero) < vector.Manhattan(closest, zero) {
			closest = p
		}
	}

	fmt.Printf("Manhattan distances:\n%+v, Manhattan distance: %f\n", closest, vector.Manhattan(zero, closest))

	// Part2: previous approach isn't very useful here.
	// Let's try a more object-oriented approach.
	// Each step (of the wire) stored in a hashed map (coordinates represent key).
	// This way querying intersections is fast and we're not abstracting the problem too much.

}
