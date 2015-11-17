// Largest palindrome product
// Problem 4
// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var max int

	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			if p := i * j; p > max && isPalindrome(p) {
				max = p
			}
		}
	}

	fmt.Println(max)
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
