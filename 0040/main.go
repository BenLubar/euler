// Champernowne's constant
// Problem 40
//
// An irrational decimal fraction is created by concatenating the positive
// integers:
//
// 0.12345678910[1]112131415161718192021...
//
// It can be seen that the 12th digit of the fractional part is 1.
//
// If d(n) represents the nth digit of the fractional part, find the value of
// the following expression.
//
// d(1) × d(10) × d(100) × d(1000) × d(10000) × d(100000) × d(1000000)
package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	r, w := io.Pipe()
	go write(w)

	var offset int64
	get := func(absolute int64) int64 {
		skip(r, absolute-offset-1)
		offset = absolute

		var buf [1]byte
		r.Read(buf[:])
		return int64(buf[0] - '0')
	}

	fmt.Println(get(1) * get(10) * get(100) * get(1000) * get(10000) * get(100000) * get(1000000))
}

func write(w io.Writer) {
	for i := 1; ; i++ {
		fmt.Fprintf(w, "%d", i)
	}
}

func skip(r io.Reader, n int64) {
	io.CopyN(ioutil.Discard, r, n)
}
