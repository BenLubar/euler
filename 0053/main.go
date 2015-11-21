// Combinatoric selections
// Problem 53
//
// There are exactly ten ways of selecting three from five, 12345:
//
// 123, 124, 125, 134, 135, 145, 234, 235, 245, and 345
//
// In combinatorics, we use the notation, 5C3 = 10.
//
// In general, nCr = (n!)/(r!(n−r)!), where r ≤ n, n! = n×(n−1)×...×3×2×1, and
// 0! = 1. It is not until n = 23, that a value exceeds one-million: 23C10 =
// 1144066.
//
// How many, not necessarily distinct, values of  nCr, for 1 ≤ n ≤ 100, are
// greater than one-million?
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var count int
	var a, b big.Int

	for n := int64(1); n <= 100; n++ {
		for r := int64(0); r <= n; r++ {
			a.MulRange(r+1, n)
			b.MulRange(1, n-r)
			a.Div(&a, &b)
			b.SetInt64(1000000)
			if a.Cmp(&b) > 0 {
				count++
			}
		}
	}

	fmt.Println(count)
}
