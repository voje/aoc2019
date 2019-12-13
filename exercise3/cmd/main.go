package main

import (
	"fmt"
	"os"
	"github.com/voje/aoc2019/exercise3/vector"
	"bufio"
	"strconv"
	"strings"
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
			prev := newPath[len(newPath) - 1]
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

	fmt.Printf("%+v", paths)


}
