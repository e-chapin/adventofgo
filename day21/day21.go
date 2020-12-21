package main

import (
	"adventofgo"
	"fmt"
	"regexp"
	"strings"
)




func main() {

	input := adventofgo.ReadFile("input.txt")

	re := regexp.MustCompile(`(.+) \(contains (.+)\)`)

	iCount := map[string]int{}
	allergens := map[string]map[string]struct{}{}

	for index, v := range input {
		fmt.Println(index, v)
		g := re.FindStringSubmatch(v)
		fmt.Println(g)

		ingredients := strings.Split(g[1], " ")
		algs := strings.Split(g[2], ", ")

		for _, i := range ingredients {
			iCount[i] += 1
		}

		for _, a := range algs {
			if _, ok := allergens[a]; !ok {
				// first time finding this allergen, build the initial set of ingredients for it.
				allergens[a] = map[string]struct{}{}
				for _, i := range ingredients {
					allergens[a][i] = struct{}{}
				}
			} else {
				// seen this allergen before, delete any ingredients that are in the existing list
				// Basically a set intersection
				for ing, _ := range allergens[a] {
					if adventofgo.IndexOf(ing, ingredients) < 0 {
						delete(allergens[a], ing)
					}
				}
			}
		}
	}

	// convert the list of sets to a single list. Doesn't matter if its not unique.
	var aList []string
	for _, a := range allergens {
		for i, _ := range a {
			aList = append(aList, i)
		}
	}

	total := 0
	// for each ingredient in master list, add its total if not found in allergens.
	for i, v := range iCount {
		if adventofgo.IndexOf(i, aList) < 0 {
			fmt.Println("Counting i", i)
			total += v
		}
	}

	fmt.Println("Day 20 part 1")
	fmt.Println(total)
	//fmt.Println("Day 20 part 2")

}
