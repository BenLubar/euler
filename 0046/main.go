// Goldbach's other conjecture
// Problem 46
//
// It was proposed by Christian Goldbach that every odd composite number can be
// written as the sum of a prime and twice a square.
//
// 9 = 7 + 2×1²
// 15 = 7 + 2×2²
// 21 = 3 + 2×3²
// 25 = 7 + 2×3²
// 27 = 19 + 2×2²
// 33 = 31 + 2×1²
//
// It turns out that the conjecture was false.
//
// What is the smallest odd composite that cannot be written as the sum of a
// prime and twice a square?
package main

import "fmt"

func main() {
	primes := []int{2}

search:
	for i := 3; ; i += 2 {
		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
			continue
		}

		for _, p := range primes {
			for n := 1; ; n++ {
				sum := p + n*n*2
				if sum > i {
					break
				}
				if sum == i {
					continue search
				}
			}
		}

		fmt.Println(i)
		return
	}
}
