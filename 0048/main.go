// Self powers
// Problem 48
//
// The series, 1¹ + 2² + 3³ + ... + 10¹⁰ = 10405071317.
//
// Find the last ten digits of the series, 1¹ + 2² + 3³ + ... + 1000¹⁰⁰⁰.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b, c big.Int

	c.SetInt64(10000000000)

	for i := int64(1); i <= 1000; i++ {
		b.SetInt64(i)
		b.Exp(&b, &b, &c)
		a.Add(&a, &b)
	}
	a.Mod(&a, &c)

	fmt.Println(a.String())
}
