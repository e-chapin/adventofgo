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
var containers = make(map[string]bool)

func findColor(color string){

	parents := containedIn[color]
	for _, c := range parents {
		containers[c.color] = true
		findColor(c.color)
	}
}

func main() {

	lines := adventofgo.ReadFile("input.txt")

	for _, line := range lines {

		colorSearch := regexp.MustCompile("(.+?) bags contain")
		bagsSearch := regexp.MustCompile("(\\d+) (.+?) bags?[,.]")
		color := colorSearch.FindAllStringSubmatch(line, -1)[0][1]
		contents := bagsSearch.FindAllStringSubmatch(line, -1)

		for _, match := range contents {
			number, _ := strconv.Atoi(match[1])
			bag := match[2]
			if _, ok := containedIn[bag]; !ok {
				containedIn[bag] = []Bag{}
			}
			containedIn[bag] = append(containedIn[bag], Bag{color, number})
		}
	}
	fmt.Println("Part 1")
	findColor("shiny gold")
	fmt.Println(len(containers))

}
