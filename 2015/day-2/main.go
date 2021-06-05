package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func total_paper(arr []string) int {
	if len(arr) < 2 {
		return 0
	}
	var t = make([]int, len(arr))
	for idx, i := range arr {
		j, _ := strconv.Atoi(i)
		t[idx] = j
	}
	sides := []int{t[0] * t[1], t[2] * t[1], t[0] * t[2]}
	sort.Ints(sides)
	return sides[0]*2 + sides[1]*2 + sides[2]*2 + sides[0]
}

func total_ribbon(arr []string) int {
	if len(arr) < 2 {
		return 0
	}
	var t = make([]int, len(arr))
	for idx, i := range arr {
		j, _ := strconv.Atoi(i)
		t[idx] = j
	}
	sort.Ints(t)
	return 2*t[0] + 2*t[1] + (t[0] * t[1] * t[2])
}

func main() {
	dat, err := ioutil.ReadFile("input")
	check(err)
	lines := strings.Split(string(dat), "\n")

	sum_paper := 0
	sum_ribbon := 0
	for _, line := range lines {
		dims := strings.Split(string(line), "x")
		sum_paper = sum_paper + total_paper(dims)
		sum_ribbon = sum_ribbon + total_ribbon(dims)
	}

	fmt.Println(sum_paper)
	fmt.Println(sum_ribbon)

}
