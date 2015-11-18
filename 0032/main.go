// Pandigital products
// Problem 32
//
// We shall say that an n-digit number is pandigital if it makes use of all the
// digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1
// through 5 pandigital.
//
// The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing
// multiplicand, multiplier, and product is 1 through 9 pandigital.
//
// Find the sum of all products whose multiplicand/multiplier/product identity
// can be written as a 1 through 9 pandigital.
//
// HINT: Some products can be obtained in more than one way so be sure to only
// include it once in your sum.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	seen := make(map[int]bool)

	/*
		1*1=1           ( 3)  9*9=81           ( 4)
		10*10=100       ( 7)  99*99=9801       ( 8)
		100*100=10000   (11)  999*999=998001   (12)
		1*10=10         ( 5)  9*99=891         ( 6)
		10*100=1000     ( 9)!!99*999=98901     (10)
		100*1000=100000 (13)  999*9999=9989001 (14)
		1*100=100       ( 7)  9*999=8991       ( 8)
		10*1000=10000   (11)  99*9999=989901   (12)
		1*1000=1000     ( 9)!!9*9999=89991     (10)
		10*10000=100000 (13)  99*99999=9899901 (14)
	*/

	for x := 10; x < 100; x++ {
		for y := 100; y < 1000; y++ {
			z := x * y
			if isPandigital(x, y, z) {
				seen[z] = true
			}
		}
	}

	for x := 1; x < 10; x++ {
		for y := 1000; y < 10000; y++ {
			z := x * y
			if isPandigital(x, y, z) {
				seen[z] = true
			}
		}
	}

	var sum int
	for i := range seen {
		sum += i
	}
	fmt.Println(sum)
}

func isPandigital(x, y, z int) bool {
	s := strconv.Itoa(x) + strconv.Itoa(y) + strconv.Itoa(z)
	if len(s) != 9 {
		return false
	}
	var have [9]bool
	for _, c := range s {
		if c == '0' {
			return false
		}
		if have[c-'1'] {
			return false
		}
		have[c-'1'] = true
	}
	return true
}
