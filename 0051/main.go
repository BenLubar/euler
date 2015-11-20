// Prime digit replacements
// Problem 51
//
// By replacing the 1st digit of the 2-digit number *3, it turns out that six
// of the nine possible values: 13, 23, 43, 53, 73, and 83, are all prime.
//
// By replacing the 3rd and 4th digits of 56**3 with the same digit, this
// 5-digit number is the first example having seven primes among the ten
// generated numbers, yielding the family: 56003, 56113, 56333, 56443, 56663,
// 56773, and 56993. Consequently 56003, being the first member of this family,
// is the smallest prime with this property.
//
// Find the smallest prime which, by replacing part of the number (not
// necessarily adjacent digits) with the same digit, is part of an eight prime
// value family.
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var primes []int
	count := make(map[string]int)

search:
	for i := 2; ; i++ {
		for _, p := range primes {
			if i%p == 0 {
				continue search
			}
		}

		primes = append(primes, i)

		s0 := strconv.Itoa(i)

		for d0 := byte('0'); d0 <= '9'; d0++ {
			for _, s1 := range replace(s0, d0) {
				count[s1]++
				if count[s1] == 8 {
					for d1 := '0'; d1 <= '2'; d1++ {
						n, err := strconv.Atoi(strings.Replace(s1, "*", string(d1), -1))
						if err != nil {
							panic(err)
						}
						if x := sort.SearchInts(primes, n); x < len(primes) && primes[x] == n {
							fmt.Println(n)
							return
						}
					}
					panic("unreachable")
				}
			}
		}
	}
}

func replace(s string, d byte) []string {
	i := strings.IndexByte(s, d)
	if i == -1 {
		return nil
	}

	result := []string{s[:i] + "*" + s[i+1:]}
	for _, after := range replace(s[i+1:], d) {
		result = append(result, s[:i]+"*"+after, s[:i+1]+after)
	}
	return result
}
