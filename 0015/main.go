// Lattice paths
// Problem 15
// Starting in the top left corner of a 2×2 grid, and only being able to move
// to the right and down, there are exactly 6 routes to the bottom right corner.
//
// How many such routes are there through a 20×20 grid?
package main

import "fmt"

func main() {
	fmt.Println(paths(20, 20))
}

var cache = make(map[[2]int]int)

func paths(w, h int) int {
	if w == 0 || h == 0 {
		return 1
	}
	if n, ok := cache[pair(w, h)]; ok {
		return n
	}

	n := paths(w-1, h) + paths(w, h-1)
	cache[pair(w, h)] = n
	return n
}

func pair(i, j int) [2]int {
	if i < j {
		i, j = j, i
	}
	return [2]int{i, j}
}
