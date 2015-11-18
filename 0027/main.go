// Quadratic primes
// Problem 27
//
// Euler discovered the remarkable quadratic formula:
//
// n² + n + 41
//
// It turns out that the formula will produce 40 primes for the consecutive
// values n = 0 to 39. However, when n = 40, 40² + 40 + 41 = 40(40 + 1) + 41 is
// divisible by 41, and certainly when n = 41, 41² + 41 + 41 is clearly
// divisible by 41.
//
// The incredible formula  n² − 79n + 1601 was discovered, which produces 80
// primes for the consecutive values n = 0 to 79. The product of the
// coefficients, −79 and 1601, is −126479.
//
// Considering quadratics of the form:
//
// n² + an + b, where |a| < 1000 and |b| < 1000
//
// where |n| is the modulus/absolute value of n
// e.g. |11| = 11 and |−4| = 4
// Find the product of the coefficients, a and b, for the quadratic expression
// that produces the maximum number of primes for consecutive values of n,
// starting with n = 0.
package main

import (
	"fmt"

	"github.com/BenLubar/memoize"
)

func main() {
	var maxA, maxB, max int

	for a := -999; a <= 999; a++ {
		for b := -999; b <= 999; b++ {
			if c := count(a, b); c > max {
				maxA, maxB, max = a, b, c
			}
		}
	}

	fmt.Println(maxA * maxB)
}

func count(a, b int) int {
	var c int
	for n := 0; isPrime(n*n + a*n + b); n++ {
		c++
	}
	return c
}

var isPrime = memoize.Memoize(func(i int) bool {
	if i < 2 {
		return false
	}
	for j := 2; j*j <= i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}).(func(int) bool)
