// Largest prime factor
// Problem 3
// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?
package main

import "fmt"

func main() {
	n := 600851475143

	for factor := 2; ; factor++ {
		for n%factor == 0 {
			if n == factor {
				fmt.Println(factor)
				return
			}

			n /= factor
		}
	}
}
