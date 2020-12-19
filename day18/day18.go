package main

import (
	"adventofgo"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	strlines := adventofgo.ReadFile("input.txt")
	total := 0
	for _, row := range strlines {
		total += run(row, regexp.MustCompile(`\([^\(\)]+\)`), evaluateExpression)
	}
	fmt.Println("Day 18 Part 1")
	fmt.Println(total)

}


func run(s string, re *regexp.Regexp, eval func(string) int) int {

	for re.MatchString(s) {
		s = re.ReplaceAllStringFunc(s, func(s string) string {
			// func(s string) in ReplaceAllStringFunc is an empty callback function to call evaluateExpression.
			// The regex "re" will find the match () non greedily and find the inner most pair and work outwards.
			return strconv.Itoa(eval(s))
		})
	}
	return eval(s)
}

func evaluateExpression (expression string) int {
	fields := strings.Fields(strings.Trim(expression, "()"))
	acc, _ := strconv.Atoi(fields[0])

	for i := 1; i < len(fields); i+=2 {
		switch n, _ := strconv.Atoi(fields[i+1]); fields[i] {
			case "+":
				acc += n
			case "*":
				acc *= n
		}
	}
	return acc
}


