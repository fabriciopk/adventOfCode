package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var lights [1000][1000]bool

const (
	on = iota + 1
	off
	toggle
)

func controlLight(x_low int, y_low int, x_up int, y_up int, ctl_type int) {
	var i, j int
	for i = x_low; i < x_up+1; i++ {
		for j = y_low; j < y_up+1; j++ {
			switch ctl_type {
			case on:
				lights[i][j] = true
			case off:
				lights[i][j] = false
			case toggle:
				lights[i][j] = !lights[i][j]
			}
		}
	}
}

func main() {

	dat, _ := ioutil.ReadFile("/home/tib/workspace/adventOfCode/2015/day-6/input")
	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		line = string(line)
		r, _ := regexp.Compile("\\d+")
		nums := r.FindAllString(line, -1)
		if len(nums) > 3 {
			x_low, _ := strconv.Atoi(nums[0])
			y_low, _ := strconv.Atoi(nums[1])
			x_up, _ := strconv.Atoi(nums[2])
			y_up, _ := strconv.Atoi(nums[3])

			if strings.Contains(line, "on") {
				controlLight(x_low, y_low, x_up, y_up, on)
			} else if strings.Contains(line, "off") {
				controlLight(x_low, y_low, x_up, y_up, off)
			} else {
				controlLight(x_low, y_low, x_up, y_up, toggle)
			}
		}
	}

	// Count lights on
	cnt := 0
	var i, j int
	for i = 0; i < 1000; i++ {
		for j = 0; j < 1000; j++ {
			if lights[i][j] {
				cnt += 1
			}
		}
	}
	fmt.Println(cnt)
}
