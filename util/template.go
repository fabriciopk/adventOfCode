package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("/home/tib/workspace/adventOfCode/2015/day-1/input")
	if err != nil {
		fmt.Println("error")
	}
	lines := string(b)
	fmt.Println(lines)
}
