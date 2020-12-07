package main

import (
	"adventofgo"
	"fmt"
	"regexp"
)

var containedIn = make(map[string][]string)
var containers = make(map[string]bool)

func findColor(color string){

	parents := containedIn[color]
	for _, c := range parents {
		containers[c] = true
		findColor(c)
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
			bag := match[2]
			if val, ok := containedIn[bag]; ok {
				containedIn[bag] = append(val, color)
			}else{
				containedIn[bag] = []string{color}
			}
		}
	}
	fmt.Println("Part 1")
	findColor("shiny gold")
	fmt.Println(len(containers))

}
