package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
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

	for _, line := range lines {

		seatId := forwardBack(line, 0, 127)
		if seatId > maxId {
			maxId = seatId
		}
	}
	fmt.Println("Part 1")
	fmt.Println(maxId)

}
