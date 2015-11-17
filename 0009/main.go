// Special Pythagorean triplet
// Problem 9
// A Pythagorean triplet is a set of three natural numbers, a < b < c, for
// which,
//
// a² + b² = c²
//
// For example, 3² + 4² = 9 + 16 = 25 = 5².
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
package main

import "fmt"

func main() {
	for c := 1; c < 1000; c++ {
		for a := 1; a < 1000-c; a++ {
			if b := 1000 - a - c; a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
