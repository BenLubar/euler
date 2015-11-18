// Digit fifth powers
// Problem 30
//
// Surprisingly there are only three numbers that can be written as the sum of
// fourth powers of their digits:
//
// 1634 = 1⁴ + 6⁴ + 3⁴ + 4⁴
// 8208 = 8⁴ + 2⁴ + 0⁴ + 8⁴
// 9474 = 9⁴ + 4⁴ + 7⁴ + 4⁴
//
// As 1 = 1⁴ is not a sum it is not included.
//
// The sum of these numbers is 1634 + 8208 + 9474 = 19316.
//
// Find the sum of all the numbers that can be written as the sum of fifth
// powers of their digits.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var pow [10]int

	for i := range pow {
		pow[i] = i * i * i * i * i
	}

	max := 1
	for i := 1; max < pow[9]*i; i, max = i+1, max*10 {
	}

	var sum int

	for i := 10; i < max; i++ {
		var n int
		for _, c := range strconv.Itoa(i) {
			n += pow[c-'0']
		}
		if i == n {
			sum += i
		}
	}

	fmt.Println(sum)
}
