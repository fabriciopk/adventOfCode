package main

import (
	"fmt"
	"io/ioutil"
)

// 2D grid of houses
type House struct {
	Lat, Long int
}

func main() {
	dat, _ := ioutil.ReadFile("/home/tib/workspace/adventOfCode/2015/day-3/input")
	str_input := string(dat)
	// House:number of presents
	var m map[House]int
	m = make(map[House]int)
	i, j := 0, 0
	m[House{i, j}] = 1
	for _, char := range str_input {
		switch char {
		case '>':
			i += 1
		case '<':
			i -= 1
		case '^':
			j += 1
		case 'v':
			j -= 1
		}
		if val, ok := m[House{i, j}]; ok {
			m[House{i, j}] = val + 1
		} else {
			m[House{i, j}] = 1
		}
	}
	fmt.Println("map:", len(m))
}
