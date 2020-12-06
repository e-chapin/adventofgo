package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)


func forwardBack(input string, min int, max int) int {

	char := input[:1]

	if char != "F" && char != "B" {
		row := min
		seat := leftRight(input, 0, 7)
		return row*8+seat
	}else{
		if char == "F" {
			newMax := (min+max)/2
			return forwardBack(input[1:], min, newMax)
		}else{
			newMin := int(math.Ceil(float64(min+max) / 2))
			return forwardBack(input[1:], newMin, max)
		}
	}

}

func leftRight(input string, min int, max int) int {

	if len(input) == 0 {
		return min
	}

	if input[:1] == "L" {
		newMax := (min+max)/2
		return leftRight(input[1:], min, newMax)
	}else{
		newMin := int(math.Ceil(float64(min+max) / 2))
		return leftRight(input[1:], newMin, max)
	}

}

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(0)
	}

	lines := strings.Split(string(content), "\n")

	maxId := -1
	var seatIds []int

	for _, line := range lines {

		seatId := forwardBack(line, 0, 127)
		seatIds = append(seatIds, seatId)
		if seatId > maxId {
			maxId = seatId
		}
	}
	fmt.Println("Part 1")
	fmt.Println(maxId)

	sort.Ints(seatIds)

	for index, seatId := range seatIds {
		if index == 0 || index == len(seatIds)-1{
			continue
		}
		next := seatIds[index+1]

		if seatId+1 != next {
			fmt.Println("Part 2")
			fmt.Println(seatId+1)
		}

	}

}
