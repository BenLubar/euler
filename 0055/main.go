// Lychrel numbers
// Problem 55
//
// If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.
//
// Not all numbers produce palindromes so quickly. For example,
//
// 349 + 943 = 1292,
// 1292 + 2921 = 4213
// 4213 + 3124 = 7337
//
// That is, 349 took three iterations to arrive at a palindrome.
//
// Although no one has proved it yet, it is thought that some numbers, like
// 196, never produce a palindrome. A number that never forms a palindrome
// through the reverse and add process is called a Lychrel number. Due to the
// theoretical nature of these numbers, and for the purpose of this problem, we
// shall assume that a number is Lychrel until proven otherwise. In addition
// you are given that for every number below ten-thousand, it will either (i)
// become a palindrome in less than fifty iterations, or, (ii) no one, with all
// the computing power that exists, has managed so far to map it to a
// palindrome. In fact, 10677 is the first number to be shown to require over
// fifty iterations before producing a palindrome: 4668731596684224866951378664
// (53 iterations, 28-digits).
//
// Surprisingly, there are palindromic numbers that are themselves Lychrel
// numbers; the first example is 4994.
//
// How many Lychrel numbers are there below ten-thousand?
//
// NOTE: Wording was modified slightly on 24 April 2007 to emphasise the
// theoretical nature of Lychrel numbers.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var count int

	var ii big.Int
	for i := 1; i < 10000; i++ {
		if isLychrel(ii.SetInt64(int64(i)), 50) {
			count++
		}
	}

	fmt.Println(count)
}

var lychrel = make(map[string]bool)

func isLychrel(i *big.Int, maxDepth int) bool {
	if is, ok := lychrel[i.String()]; ok {
		return is
	}

	n := addReverse(i)
	if isPalindrome(n) {
		lychrel[i.String()] = false
		return false
	}

	if maxDepth <= 0 {
		lychrel[i.String()] = true
		return true
	}

	is := isLychrel(n, maxDepth-1)
	lychrel[i.String()] = is
	return is
}

var ten = big.NewInt(10)

func addReverse(i *big.Int) *big.Int {
	j, k := big.NewInt(1), big.NewInt(1)
	for i.Cmp(j) >= 0 {
		j.Mul(j, ten)
	}
	j.Div(j, ten)

	var n, tmp big.Int
	n.Set(i)
	for j.Sign() != 0 {
		tmp.Div(i, j)
		tmp.Mod(&tmp, ten)
		tmp.Mul(&tmp, k)
		n.Add(&n, &tmp)
		j.Div(j, ten)
		k.Mul(k, ten)
	}
	return &n
}

func isPalindrome(i *big.Int) bool {
	s := i.String()

	for j, k := 0, len(s)-1; j < k; j, k = j+1, k-1 {
		if s[j] != s[k] {
			return false
		}
	}

	return true
}
