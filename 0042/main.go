// Coded triangle numbers
// Problem 42
//
// The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1);
// so the first ten triangle numbers are:
//
// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
//
// By converting each letter in a word to a number corresponding to its
// alphabetical position and adding these values we form a word value. For
// example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word
// value is a triangle number then we shall call the word a triangle word.
//
// Using words.txt (right click and 'Save Link/Target As...'), a 16K text file
// containing nearly two-thousand common English words, how many are triangle
// words?
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	var isTriangle [15 * 26]bool

	for i, j := 2, 1; j < len(isTriangle); i, j = i+1, j+i {
		isTriangle[j] = true
	}

	words, err := getWords()
	if err != nil {
		panic(err)
	}

	var count int

	for _, w := range words {
		var value int
		for _, c := range w {
			value += int(c - 'A' + 1)
		}
		if isTriangle[value] {
			count++
		}
	}

	fmt.Println(count)
}

func getWords() ([]string, error) {
	resp, err := http.Get("https://projecteuler.net/project/resources/p042_words.txt")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var words []string
	err = json.NewDecoder(io.MultiReader(strings.NewReader("["), resp.Body, strings.NewReader("]"))).Decode(&words)
	return words, err
}
