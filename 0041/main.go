// Pandigital prime
// Problem 41
//
// We shall say that an n-digit number is pandigital if it makes use of all the
// digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is
// also prime.
//
// What is the largest n-digit pandigital prime that exists?
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// no 9-digit pandigital number can be prime because the digits will
	// add up to 45, which means it's divisible by 9.
	var nonPrime [1e8]bool

	nonPrime[0] = true
	nonPrime[1] = true

	var max int

	for i, skip := range &nonPrime {
		if skip {
			continue
		}

		for j := i + i; j < len(nonPrime); j += i {
			nonPrime[j] = true
		}

		if isPandigital(i) {
			max = i
		}
	}

	fmt.Println(max)
}

func isPandigital(n int) bool {
	s := strconv.Itoa(n)

	var have [10]bool
	have[0] = true

	for _, c := range s {
		if len(s) < int(c-'0') {
			return false
		}
		if have[c-'0'] {
			return false
		}
		have[c-'0'] = true
	}
	return true
}
