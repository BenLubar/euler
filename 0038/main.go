// Pandigital multiples
// Problem 38
//
// Take the number 192 and multiply it by each of 1, 2, and 3:
//
// 192 × 1 = 192
// 192 × 2 = 384
// 192 × 3 = 576
//
// By concatenating each product we get the 1 to 9 pandigital, 192384576. We
// will call 192384576 the concatenated product of 192 and (1,2,3)
//
// The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4,
// and 5, giving the pandigital, 918273645, which is the concatenated product
// of 9 and (1,2,3,4,5).
//
// What is the largest 1 to 9 pandigital 9-digit number that can be formed as
// the concatenated product of an integer with (1,2, ... , n) where n > 1?
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var max string

	for n := 2; n <= 9; n++ {
		for x := 1; ; x++ {
			s := concatenatedProduct(x, n)
			if len(s) > 9 {
				break
			}

			if s > max && isPandigital(s) {
				max = s
			}
		}
	}

	fmt.Println(max)
}

func concatenatedProduct(x, n int) string {
	var b []byte

	for i := 1; i <= n; i++ {
		b = strconv.AppendInt(b, int64(x*i), 10)
	}

	return string(b)
}

func isPandigital(s string) bool {
	if len(s) != 9 {
		return false
	}
	var have [9]bool
	for _, c := range s {
		if c == '0' {
			return false
		}
		if have[c-'1'] {
			return false
		}
		have[c-'1'] = true
	}
	return true
}
