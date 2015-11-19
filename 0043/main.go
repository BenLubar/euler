// Sub-string divisibility
// Problem 43
//
// The number, 1406357289, is a 0 to 9 pandigital number because it is made up
// of each of the digits 0 to 9 in some order, but it also has a rather
// interesting sub-string divisibility property.
//
// Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we
// note the following:
//
// d2d3d4=406 is divisible by 2
// d3d4d5=063 is divisible by 3
// d4d5d6=635 is divisible by 5
// d5d6d7=357 is divisible by 7
// d6d7d8=572 is divisible by 11
// d7d8d9=728 is divisible by 13
// d8d9d10=289 is divisible by 17
//
// Find the sum of all 0 to 9 pandigital numbers with this property.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := []byte("0123456789")

	var sum int

	for {
		permute(n)
		s := string(n)
		if s == "0123456789" {
			break
		}

		if s[0] != '0' &&
			check(s[1:4], 2) &&
			check(s[2:5], 3) &&
			check(s[3:6], 5) &&
			check(s[4:7], 7) &&
			check(s[5:8], 11) &&
			check(s[6:9], 13) &&
			check(s[7:10], 17) {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			sum += i
		}
	}

	fmt.Println(sum)
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

func check(s string, n int64) bool {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i%n == 0
}
