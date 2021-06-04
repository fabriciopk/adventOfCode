package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func first(input string) {
	up := strings.Count(input, "(")
	down := strings.Count(input, ")")
	fmt.Println("What floor do the instructions take Santa:", up-down)
}

func second(input string) {
	cnt := 0
	for pos, char := range input {
		if char == '(' {
			cnt = cnt + 1
		} else {
			cnt = cnt - 1
		}
		if cnt == -1 {
			fmt.Println("What is the position of the basement character: ", pos+1)
			break
		}
	}
}

func main() {
	dat, err := ioutil.ReadFile("input")
	check(err)
	str_input := string(dat)
	first(str_input)
	second(str_input)
}
