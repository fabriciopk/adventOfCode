package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var lights [1000][1000]int

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
				lights[i][j] += 1
			case off:
				val := lights[i][j] - 1
				if val < 0 {
					lights[i][j] = 0
				} else {
					lights[i][j] = val
				}
			case toggle:
				lights[i][j] += 2
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
			cnt += lights[i][j]
		}
	}
	fmt.Println(cnt)
}
