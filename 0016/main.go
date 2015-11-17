// Power digit sum
// Problem 16
// 2¹⁵ = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
//
// What is the sum of the digits of the number 2¹⁰⁰⁰?
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n big.Int

	n.SetInt64(1)
	n.Lsh(&n, 1000)

	var sum int
	for _, c := range n.String() {
		sum += int(c - '0')
	}

	fmt.Println(sum)
}
