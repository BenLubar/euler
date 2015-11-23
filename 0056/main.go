// Powerful digit sum
// Problem 56
//
// A googol (10¹⁰⁰) is a massive number: one followed by one-hundred zeros;
// 100¹⁰⁰ is almost unimaginably large: one followed by two-hundred zeros.
// Despite their size, the sum of the digits in each number is only 1.
//
// Considering natural numbers of the form, aᵇ, where a, b < 100, what is the
// maximum digital sum?
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var bigA, bigB, z big.Int

	var max int

	for a := 1; a < 100; a++ {
		bigA.SetInt64(int64(a))
		for b := 1; b < 100; b++ {
			bigB.SetInt64(int64(b))
			sum := digitalSum(z.Exp(&bigA, &bigB, nil).String())
			if sum > max {
				max = sum
			}
		}
	}

	fmt.Println(max)
}

func digitalSum(s string) int {
	var n int

	for _, c := range s {
		n += int(c - '0')
	}

	return n
}
