package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	
	spokenNumbers := make(map[int][]int)

	input := "8,11,0,19,1,2"
	listIn := strings.Split(input, ",")
	startLen := len(listIn)+1

	var lastSpoken int
	lastSpoken = 0

	for i, c := range listIn {
		ival, _ := strconv.Atoi(c)
		spokenNumbers[ival] = []int{i+1}
		lastSpoken = ival
	}

	var next int
	for i := startLen; i <= 2020; i++ {
		var timesSpoken = len(spokenNumbers[lastSpoken])
		if timesSpoken == 1{
			next = 0
		} else {
			next = spokenNumbers[lastSpoken][timesSpoken-1] - spokenNumbers[lastSpoken][timesSpoken-2]
		}

		if _, ok := spokenNumbers[next]; ok {
			spokenNumbers[next] = append(spokenNumbers[next], i)
		} else {
			spokenNumbers[next] = []int{i}
		}
		lastSpoken = next
	}
	fmt.Println("Day 15 Part 1")
	fmt.Println(lastSpoken)

	//fmt.Println(spokenNumbers)

}
