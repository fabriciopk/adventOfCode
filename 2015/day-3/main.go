package main

import (
	"fmt"
	"io/ioutil"
)

// 2D grid of houses
type House struct {
	Lat, Long int
}

func first_solution(input string) int {
	var m map[House]int
	m = make(map[House]int)
	i, j := 0, 0
	m[House{i, j}] = 1
	for _, char := range input {
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
	return len(m)
}

func move(char rune, h House) House {
	m := h
	switch char {
	case '>':
		m.Lat += 1
	case '<':
		m.Lat -= 1
	case '^':
		m.Long += 1
	case 'v':
		m.Long -= 1
	}
	return m
}

func second_solution(input string) int {

	var houses map[House]int
	houses = make(map[House]int)

	santaHouse := House{0, 0}
	robotHouse := House{0, 0}
	houses[santaHouse] = 1
	var m House
	for index, char := range input {
		if index%2 == 0 {
			m = move(char, santaHouse)
			santaHouse = m
		} else {
			m = move(char, robotHouse)
			robotHouse = m
		}
		if val, ok := houses[House{m.Lat, m.Long}]; ok {
			houses[House{m.Lat, m.Long}] = val + 1
		} else {
			houses[House{m.Lat, m.Long}] = 1
		}

	}

	return len(houses)
}

func main() {
	dat, _ := ioutil.ReadFile("/home/tib/workspace/adventOfCode/2015/day-3/input")
	str_input := string(dat)

	fmt.Println("Houses that received at least one present:", first_solution(str_input))
	fmt.Println("Houses that received at least one present w/ robot:", second_solution(str_input))

}
