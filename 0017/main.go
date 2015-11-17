// Number letter counts
// Problem 17
// If the numbers 1 to 5 are written out in words: one, two, three, four, five,
// then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//
// If all the numbers from 1 to 1000 (one thousand) inclusive were written out
// in words, how many letters would be used?
//
// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
// forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20
// letters. The use of "and" when writing out numbers is in compliance with
// British usage.
package main

import "fmt"

func main() {
	var sum int

	for i := 1; i <= 1000; i++ {
		sum += compute(i)
	}

	fmt.Println(sum)
}

func compute(i int) int {
	if i < 20 {
		return [20]int{
			len("zero"),
			len("one"),
			len("two"),
			len("three"),
			len("four"),
			len("five"),
			len("six"),
			len("seven"),
			len("eight"),
			len("nine"),
			len("ten"),
			len("eleven"),
			len("twelve"),
			len("thirteen"),
			len("fourteen"),
			len("fifteen"),
			len("sixteen"),
			len("seventeen"),
			len("eighteen"),
			len("nineteen"),
		}[i]
	}
	if i < 100 {
		if rem := i % 10; rem != 0 {
			return compute(i-rem) + compute(rem)
		}
		return [10]int{
			0,
			0,
			len("twenty"),
			len("thirty"),
			len("forty"),
			len("fifty"),
			len("sixty"),
			len("seventy"),
			len("eighty"),
			len("ninety"),
		}[i/10]
	}

	if i < 1000 {
		if rem := i % 100; rem != 0 {
			return compute(i-rem) + len("and") + compute(rem)
		}
		return compute(i/100) + len("hundred")
	}

	if i == 1000 {
		return len("one") + len("thousand")
	}

	panic("unreachable")
}
