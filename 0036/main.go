// Double-base palindromes
// Problem 36
//
// The decimal number, 585 = 1001001001₂ (binary), is palindromic in both bases.
//
// Find the sum of all numbers, less than one million, which are palindromic in
// base 10 and base 2.
//
// (Please note that the palindromic number, in either base, may not include
// leading zeros.)
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum int

	for i := 1; i < 1000000; i++ {
		if isPalindrome(i, 2) && isPalindrome(i, 10) {
			sum += i
		}
	}

	fmt.Println(sum)
}

func isPalindrome(n, base int) bool {
	s := strconv.FormatInt(int64(n), base)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
