// Names scores
// Problem 22
// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file
// containing over five-thousand first names, begin by sorting it into
// alphabetical order. Then working out the alphabetical value for each name,
// multiply this value by its alphabetical position in the list to obtain a
// name score.
//
// For example, when the list is sorted into alphabetical order, COLIN, which
// is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN
// would obtain a score of 938 × 53 = 49714.
//
// What is the total of all the name scores in the file?
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

func main() {
	names, err := getNames()
	if err != nil {
		panic(err)
	}

	sort.Strings(names)

	var total int

	for i, name := range names {
		sum := len(name)
		for _, c := range name {
			sum += int(c - 'A')
		}

		total += (i + 1) * sum
	}

	fmt.Println(total)
}

func getNames() ([]string, error) {
	resp, err := http.Get("https://projecteuler.net/project/resources/p022_names.txt")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP request returned status " + resp.Status)
	}

	var lines []string
	err = json.NewDecoder(io.MultiReader(strings.NewReader("["), resp.Body, strings.NewReader("]"))).Decode(&lines)
	return lines, err
}
