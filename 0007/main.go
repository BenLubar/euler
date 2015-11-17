// 10001st prime
// Problem 7
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
// that the 6th prime is 13.
//
// What is the 10 001st prime number?
package main

import "fmt"

func main() {
	ch := make(chan int)
	go gen(ch)

	for n := 1; ; n++ {
		prime := <-ch
		if n == 10001 {
			fmt.Println(prime)
			return
		}
		ch1 := make(chan int)
		go sieve(ch1, ch, prime)
		ch = ch1
	}
}

func gen(out chan<- int) {
	for i := 2; ; i++ {
		out <- i
	}
}

func sieve(out chan<- int, in <-chan int, prime int) {
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
}
