// Summation of primes
// Problem 10
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
package main

import "fmt"

func main() {
	var numbers [2000000]int
	for i := range numbers {
		if i < 2 {
			continue
		}
		numbers[i] = i
	}

	var sum int

	for _, n := range &numbers {
		if n == 0 {
			continue
		}

		for i := n + n; i < len(numbers); i += n {
			numbers[i] = 0
		}

		sum += n
	}

	fmt.Println(sum)
}
