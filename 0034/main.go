// Digit factorials
// Problem 34
//
// 145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
//
// Find the sum of all numbers which are equal to the sum of the factorial of
// their digits.
//
// Note: as 1! = 1 and 2! = 2 are not sums they are not included.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var fact [10]int

	for i := range fact {
		if i == 0 {
			fact[0] = 1
		} else {
			fact[i] = fact[i-1] * i
		}
	}

	max := 10
	for i := 1; max < i*fact[9]; i, max = i+1, max*10 {
	}

	var sum int

	for i := 10; i < max; i++ {
		var n int
		for _, c := range strconv.Itoa(i) {
			n += fact[c-'0']
		}
		if i == n {
			sum += i
		}
	}

	fmt.Println(sum)
}
