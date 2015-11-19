// Prime permutations
// Problem 49
//
// The arithmetic sequence, 1487, 4817, 8147, in which each of the terms
// increases by 3330, is unusual in two ways: (i) each of the three terms are
// prime, and, (ii) each of the 4-digit numbers are permutations of one
// another.
//
// There are no arithmetic sequences made up of three 1-, 2-, or 3-digit
// primes, exhibiting this property, but there is one other 4-digit increasing
// sequence.
//
// What 12-digit number do you form by concatenating the three terms in this
// sequence?
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var nonPrime [10000]bool
	nonPrime[0] = true
	nonPrime[1] = true

	for i, skip := range &nonPrime {
		if skip {
			continue
		}

		for j := i + i; j < len(nonPrime); j += i {
			nonPrime[j] = true
		}
	}

	for i, skip := range nonPrime {
		if skip || i < 1000 || i == 1487 {
			continue
		}

		s0, s1, s2 := strconv.Itoa(i), "", ""

		b := []byte(s0)

		for j := 0; j < 4*3*2*1; j++ {
			permute(b)
			s := string(b)
			if s == s0 || s == s1 || s[0] == '0' {
				continue
			}
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			if nonPrime[n] {
				continue
			}

			if s1 != "" {
				s2 = s
				break
			}
			s1 = s
		}

		if s2 != "" {
			s := []string{s0, s1, s2}
			sort.Strings(s)
			i0, err := strconv.Atoi(s[0])
			if err != nil {
				panic(err)
			}
			i1, err := strconv.Atoi(s[1])
			if err != nil {
				panic(err)
			}
			i2, err := strconv.Atoi(s[2])
			if err != nil {
				panic(err)
			}
			if i0-i1 != i1-i2 {
				continue
			}
			fmt.Println(s[0] + s[1] + s[2])
			return
		}
	}
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
