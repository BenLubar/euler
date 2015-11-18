// Lexicographic permutations
// Problem 24
// A permutation is an ordered arrangement of objects. For example, 3124 is one
// possible permutation of the digits 1, 2, 3 and 4. If all of the permutations
// are listed numerically or alphabetically, we call it lexicographic order.
// The lexicographic permutations of 0, 1 and 2 are:
//
// 012   021   102   120   201   210
//
// What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4,
// 5, 6, 7, 8 and 9?
package main

import "fmt"

func main() {
	x := []byte("0123456789")
	for i := 1; i < 1000000; i++ {
		permute(x)
	}
	fmt.Println(string(x))
}

func permute(b []byte) {
	for i := len(b) - 1; i > 0; i-- {
		if b[i-1] < b[i] {
			for j := len(b) - 1; j >= i; j-- {
				if b[i-1] < b[j] {
					b[i-1], b[j] = b[j], b[i-1]
					reverse(b[i:])
					return
				}
			}
		}
	}

	reverse(b)
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
