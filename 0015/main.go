// Lattice paths
// Problem 15
// Starting in the top left corner of a 2×2 grid, and only being able to move
// to the right and down, there are exactly 6 routes to the bottom right corner.
//
// How many such routes are there through a 20×20 grid?
package main

import (
	"fmt"

	"github.com/BenLubar/memoize"
)

func main() {
	fmt.Println(paths(20, 20))
}

var paths func(w, h int) int

func init() {
	paths = memoize.Memoize(func(w, h int) int {
		if w == 0 || h == 0 {
			return 1
		}

		n := paths(w-1, h) + paths(w, h-1)
		return n
	}).(func(w, h int) int)
}
