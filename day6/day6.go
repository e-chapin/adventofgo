package main

import (
	"adventofgo"
	"fmt"
)

func main() {

	lines := adventofgo.ReadFile("input.txt")

	sum := 0

	answers := make(map[string]bool)

	for index, line := range lines {
		if line == ""{
			sum += len(answers)
			answers = make(map[string]bool)
		}

		for _, char := range line {
			answers[string(char)] = true
		}

		if index == len(lines)-1 {
			sum += len(answers)
		}
	}

	fmt.Println("Part 1")
	fmt.Println(sum)

	
}
