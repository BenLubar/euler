// Permuted multiples
// Problem 52
//
// It can be seen that the number, 125874, and its double, 251748, contain
// exactly the same digits, but in a different order.
//
// Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x,
// contain the same digits.
package main

import (
	"fmt"
	"strconv"
)

func main() {
search:
	for i := 1; ; i++ {
		var d1 [10]int
		for _, c := range strconv.Itoa(i) {
			d1[c-'0']++
		}

		for n := 2; n <= 6; n++ {
			var dn [10]int
			for _, c := range strconv.Itoa(i * n) {
				dn[c-'0']++
			}
			if d1 != dn {
				continue search
			}
		}

		fmt.Println(i)
		return
	}
}
