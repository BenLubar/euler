// Digit cancelling fractions
// Problem 33
//
// The fraction 49/98 is a curious fraction, as an inexperienced mathematician
// in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which
// is correct, is obtained by cancelling the 9s.
//
// We shall consider fractions like, 30/50 = 3/5, to be trivial examples.
//
// There are exactly four non-trivial examples of this type of fraction, less
// than one in value, and containing two digits in the numerator and
// denominator.
//
// If the product of these four fractions is given in its lowest common terms,
// find the value of the denominator.
package main

import "fmt"

func main() {
	n, d := 1, 1

	for digit := 1; digit < 10; digit++ {
		for a := 1; a < 10; a++ {
			for b := 1; b < 10; b++ {
				if a == b {
					continue
				}

				n0 := a*10 + digit
				d0 := digit*10 + b
				n1 := digit*10 + a
				d1 := b*10 + digit

				if n0 < d0 && n0*b == d0*a {
					n *= n0
					d *= d0
				}
				if n1 < d1 && n1*b == d1*a {
					n *= n1
					d *= d1
				}
			}
		}
	}

	fmt.Println(d / gcd(n, d))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
