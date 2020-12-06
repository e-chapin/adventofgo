package main

import (
	"adventofgo"
	"fmt"
)

func checkAnswers(personCount int, answers map[string]int) int {

	count := 0
	for _, letter := range answers {
		if letter == personCount {
			count += 1
		}
	}
	return count
}

func main() {

	lines := adventofgo.ReadFile("input.txt")

	anySum, everySum, personCount := 0, 0, 0

	answers := make(map[string]int)

	for index, line := range lines {
		if line == "" {
			anySum += len(answers)
			everySum += checkAnswers(personCount, answers)
			answers = make(map[string]int)
			personCount = 0
			continue
		}

		personCount += 1

		for _, char := range line {
			s := string(char)
			answers[s] = answers[s]+1
		}

		if index == len(lines)-1 {
			anySum += len(answers)
			everySum += checkAnswers(personCount, answers)
		}
	}

	fmt.Println("Part 1")
	fmt.Println(anySum)

	fmt.Println("Part 2")
	fmt.Println(everySum)

	
}
