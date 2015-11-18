// Coin sums
// Problem 31
//
// In England the currency is made up of pound, £, and pence, p, and there are
// eight coins in general circulation:
//
// 1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
//
// It is possible to make £2 in the following way:
//
// 1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
//
// How many different ways can £2 be made using any number of coins?
package main

import (
	"fmt"

	"github.com/BenLubar/memoize"
)

func main() {
	fmt.Println(ways(200, 200))
}

var ways func(int, int) int

func init() {
	ways = memoize.Memoize(func(p, max int) int {
		if p == 0 {
			return 1
		}

		var total int
		if p >= 1 && max >= 1 {
			total += ways(p-1, 1)
		}
		if p >= 2 && max >= 2 {
			total += ways(p-2, 2)
		}
		if p >= 5 && max >= 5 {
			total += ways(p-5, 5)
		}
		if p >= 10 && max >= 10 {
			total += ways(p-10, 10)
		}
		if p >= 20 && max >= 20 {
			total += ways(p-20, 20)
		}
		if p >= 50 && max >= 50 {
			total += ways(p-50, 50)
		}
		if p >= 100 && max >= 100 {
			total += ways(p-100, 100)
		}
		if p >= 200 && max >= 200 {
			total += ways(p-200, 200)
		}
		return total
	}).(func(int, int) int)
}
