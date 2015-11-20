// Consecutive prime sum
// Problem 50
//
// The prime 41, can be written as the sum of six consecutive primes:
//
// 41 = 2 + 3 + 5 + 7 + 11 + 13
//
// This is the longest sum of consecutive primes that adds to a prime below
// one-hundred.
//
// The longest sum of consecutive primes below one-thousand that adds to a
// prime, contains 21 terms, and is equal to 953.
//
// Which prime, below one-million, can be written as the sum of the most
// consecutive primes?
package main

import "fmt"

func main() {
	var nonPrime [1000000]bool

	nonPrime[0] = true
	nonPrime[1] = true

	for i, skip := range &nonPrime {
		if skip {
			continue
		}

		for j := i + i; j < len(nonPrime); j += i {
			nonPrime[j] = true
		}
	}

	var best, max int

	for i, skip := range &nonPrime {
		if skip {
			continue
		}

		for j := 0; j < i; j++ {
			if nonPrime[j] {
				continue
			}

			var sum, count int

			for k := j; sum < i; k++ {
				if nonPrime[k] {
					continue
				}

				sum += k
				count++
			}

			if sum == i && max < count {
				best = i
				max = count
			}
		}
	}

	fmt.Println(best)
}
