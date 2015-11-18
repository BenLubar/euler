// Number spiral diagonals
// Problem 28
//
// Starting with the number 1 and moving to the right in a clockwise direction
// a 5 by 5 spiral is formed as follows:
//
// [21]22 23 24[25]
//  20[ 7] 8[ 9]10
//  19  6[ 1] 2 11
//  18[ 5]  4[ 3]12
// [17] 16 15 14[13]
//
// It can be verified that the sum of the numbers on the diagonals is 101.
//
// What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral
// formed in the same way?
package main

import "fmt"

func main() {
	var spiral [1001][1001]int

	i := 0
	x, y := len(spiral)/2, len(spiral)/2
	for j := -1; j < len(spiral); {
		for k := 0; k < j; k++ {
			i++
			spiral[x][y] = i
			y++
		}
		j++
		for k := 0; k < j; k++ {
			i++
			spiral[x][y] = i
			x--
		}
		for k := 0; k < j; k++ {
			i++
			spiral[x][y] = i
			y--
		}
		j++
		for k := 0; k < j; k++ {
			i++
			spiral[x][y] = i
			x++
		}
	}

	var sum int
	sum -= 1
	for i := 0; i < len(spiral); i++ {
		sum += spiral[i][i]
		sum += spiral[len(spiral)-1-i][i]
	}

	fmt.Println(sum)
}
