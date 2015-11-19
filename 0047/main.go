package main

import "fmt"

func main() {
	var primes []int

	var count int

	for i := 2; ; i++ {
		d, n := 0, i
		for _, p := range primes {
			if n%p == 0 {
				n /= p
				d++
				for n%p == 0 {
					n /= p
				}
			}
		}

		if d == 0 {
			count = 0
			primes = append(primes, i)
		} else if d == 4 {
			count++
			if count == 4 {
				fmt.Println(i - 4 + 1)
				return
			}
		} else {
			count = 0
		}
	}
}
