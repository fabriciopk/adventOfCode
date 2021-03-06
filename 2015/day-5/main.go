package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func niceString(line string) bool {

	bad_strings := map[string]bool{"ab": true, "cd": true, "pq": true, "xy": true}
	good_vowels := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true}
	nice_twice := false
	good_vowels_cnt := 0

	// Contains at least one letter twice in a row
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			nice_twice = true
		} else if bad_strings[line[i:i+2]] {
			// has a bad string
			return false
		}
	}
	for _, r := range line {
		if good_vowels[string(r)] {
			good_vowels_cnt += 1
		}
	}
	return nice_twice && good_vowels_cnt >= 3
}

func niceString2(line string) bool {
	nice_pair := false
	nice_overlap := false
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			nice_pair = true
		}
		if strings.Count(line, line[i:i+2]) >= 2 {
			nice_overlap = true
		}
	}
	return nice_pair && nice_overlap
}

func main() {
	dat, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(dat), "\n")

	sum_nice := 0
	for _, line := range lines {
		if niceString2(line) {
			sum_nice = sum_nice + 1
		}
	}

	fmt.Println(sum_nice)

}
