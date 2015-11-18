// Integer right triangles
// Problem 39
//
// If p is the perimeter of a right angle triangle with integral length sides,
// {a,b,c}, there are exactly three solutions for p = 120.
//
// {20,48,52}, {24,45,51}, {30,40,50}
//
// For which value of p ≤ 1000, is the number of solutions maximised?
package main

import "fmt"

func main() {
	var maxP, max int

	for p := 1; p <= 1000; p++ {
		var count int

		for a := 1; a+a < p; a++ {
			for b := 1; b+a < p; b++ {
				if c := p - a - b; a*a+b*b == c*c {
					count++
				}
			}
		}

		if count > max {
			maxP, max = p, count
		}
	}

	fmt.Println(maxP)
}
