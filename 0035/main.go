// Circular primes
// Problem 35
//
// The number, 197, is called a circular prime because all rotations of the
// digits: 197, 971, and 719, are themselves prime.
//
// There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37,
// 71, 73, 79, and 97.
//
// How many circular primes are there below one million?
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var notPrime [1000000]bool

	notPrime[0] = true
	notPrime[1] = true

	for i, skip := range &notPrime {
		if skip {
			continue
		}

		for j := i + i; j < len(notPrime); j += i {
			notPrime[j] = true
		}
	}

	var count int

search:
	for i, skip := range &notPrime {
		if skip {
			continue
		}

		s := strconv.Itoa(i)
		for j := range s {
			n, err := strconv.Atoi(s[j:] + s[:j])
			if err != nil {
				panic(err)
			}
			if notPrime[n] {
				continue search
			}
		}

		count++
	}

	fmt.Println(count)
}
