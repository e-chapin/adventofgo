package main

import (
	"adventofgo"
	"fmt"
	"regexp"
	"strings"
)

func createRuleOptions(index string, rules map[string]string) []string {
	// could probably generate a regex string more quickly, but this was easier to think through for me.
	ruleVal := rules[index]
	if ruleVal == `"a"` {
		return []string{"a"}
	}
	if ruleVal == `"b"` {
		return []string{"b"}
	}

	options := strings.Split(ruleVal, " | ")
	// there is always at least one set of rules
	firstSet := strings.Split(options[0],  " ")

	// Get the first rule
	var newRulesOne = createRuleOptions(firstSet[0], rules)
	// Get the rest of the rules
	for i := 1; i < len(firstSet); i++ {
		subRuleIndexI := createRuleOptions(firstSet[i], rules)
		var newNewRulesOne []string
		for _, y := range subRuleIndexI {
			for _, x := range newRulesOne {
				// Post-pend each of the new rule variations onto the existing list.
				// so rule x in the existing list will be duplicated with and appended with new rule y.
				newNewRulesOne = append(newNewRulesOne, x + y)
			}
		}
		// replace existing list with new list
		newRulesOne = newNewRulesOne
	}

	if len(options) == 1 {
		return newRulesOne
	}

	// do the same for second list if exists. Never more than two.
	secondSet := strings.Split(options[1],  " ")
	newRulesTwo := createRuleOptions(secondSet[0], rules)
	for i := 1; i < len(secondSet); i++ {
		subRuleIndexI := createRuleOptions(secondSet[i], rules)
		var newNewRulesTwo []string
		for _, y := range subRuleIndexI {
			for _, x := range newRulesTwo {
				newNewRulesTwo = append(newNewRulesTwo, x + y)
			}
		}
		newRulesTwo = newNewRulesTwo
	}

	return append(newRulesOne, newRulesTwo...)

}

func main() {

	var rules = make(map[string]string)

	strlines := adventofgo.ReadFile("input.txt")
	index := 0

	for {
		line := strlines[index]
		if line == "" {
			index += 1
			break
		}
		r := regexp.MustCompile(`([\d]+): ([0-9  \|"a""b"]+)`).FindAllStringSubmatch(line, -1)[0]
		rules[r[1]] = r[2]
		index += 1
	}

	count := 0
	validMessages := createRuleOptions("0", rules)
	for index := index; index < len(strlines); index += 1 {
		message := strlines[index]
		if adventofgo.IndexOf(message, validMessages) > 0 {
			count += 1
		}
	}
	fmt.Println("Day 19 Part 1")
	fmt.Println(count)
}
