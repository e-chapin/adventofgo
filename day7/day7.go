package main

import (
	"adventofgo"
	"fmt"
	"regexp"
	"strconv"
)

type Bag struct {
	color string
	number int

}

var containedIn = make(map[string][]Bag)
var contains = make(map[string][]Bag)
var containers = make(map[string]bool)

func findColor(color string) {

	parents := containedIn[color]
	for _, c := range parents {
		containers[c.color] = true
		findColor(c.color)
	}
}

func calculateInnerBags(color string) int {

	total := 0
	for _, innerBag := range contains[color] {
		total += innerBag.number + innerBag.number * calculateInnerBags(innerBag.color)
	}
	return total
}


func main() {

	lines := adventofgo.ReadFile("input.txt")

	for _, line := range lines {

		color := regexp.MustCompile("(.+?) bags contain").FindAllStringSubmatch(line, -1)[0][1]
		contents := regexp.MustCompile("(\\d+) (.+?) bags?[,.]").FindAllStringSubmatch(line, -1)

		for _, match := range contents {
			number, _ := strconv.Atoi(match[1])
			bag := match[2]
			containedIn[bag] = append(containedIn[bag], Bag{color, number})
			contains[color] = append(contains[color], Bag{bag, number})
		}
	}

	fmt.Println("Part 1")
	findColor("shiny gold")
	fmt.Println(len(containers))

	fmt.Println("Part 2")
	fmt.Println(calculateInnerBags("shiny gold"))

}
