// Truncatable primes
// Problem 37
//
// The number 3797 has an interesting property. Being prime itself, it is
// possible to continuously remove digits from left to right, and remain prime
// at each stage: 3797, 797, 97, and 7. Similarly we can work from right to
// left: 3797, 379, 37, and 3.
//
// Find the sum of the only eleven primes that are both truncatable from left
// to right and right to left.
//
// NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
package main

import "fmt"

func main() {
	isPrime := make(map[int]bool)

	var sum, count int

search:
	for prime := 2; count < 11; prime++ {
		for p := range isPrime {
			if prime%p == 0 {
				continue search
			}
		}

		isPrime[prime] = true

		if prime < 10 {
			continue
		}

		n := prime
		for i := 10; i < prime; i *= 10 {
			n /= 10
			if !isPrime[n] || !isPrime[prime%i] {
				continue search
			}
		}

		sum += prime
		count++
	}

	fmt.Println(sum)
}
