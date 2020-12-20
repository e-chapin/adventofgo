package main

import (
	"adventofgo"
	"fmt"
	"regexp"
	"strings"
)

func main() {

	input := adventofgo.ReadFile("input.txt")
	fmt.Println("Day 19 part 1")
	fmt.Println(partOne(input))
	fmt.Println("Day 19 part 2")
	fmt.Println(partTwo(input))
}

func partOne(input []string) int {

	rules, messages := parseInput(input)
	re := regexp.MustCompile(`\d+`)
	for {
		done := true
		for k := range rules {
			rules[k] = re.ReplaceAllStringFunc(rules[k], func (s string) string {
				// convert s (digits) to a regex string based on the rules.
				// the outermost for (while) loop will keep running until every rule has propagated down to an
				// (a) or (b)
				done = false
				i := adventofgo.AsInt(s)
				return rules[i]
			})
		}
		if done {
			break
		}
	}


	regexRules := map[int]*regexp.Regexp{}

	for k := range rules {
		replacer := strings.NewReplacer("\"", "", " ", "")
		regexRules[k] = regexp.MustCompile("^" + replacer.Replace(rules[k]) + "$")
	}

	first := regexRules[0]

	count := 0
	for _, m := range messages {
		if first.MatchString(m) {
			count += 1
		}
	}
	return count

}

func partTwo(input []string) int {

	rules, messages := parseInput(input)
	re := regexp.MustCompile(`\d+`)

	rules[8] = "(42 | 42 8)"
	rules[11] = "(42 31 | 42 11 31)"

	rule8 := make(map[int]int)
	rule11 := make(map[int]int)
	max := 10

	for {
		done := true
		for k := range rules {
			rules[k] = re.ReplaceAllStringFunc(rules[k], func (s string) string {
				// convert s (digits) to a regex string based on the rules.
				// the outermost for (while) loop will keep running until every rule has propagated down to an
				// (a) or (b)

				done = false
				i := adventofgo.AsInt(s)

				// keep track of how many times this particular index has looped. Cut it off when its done enough times.
				if i == 8 {
					if _, ok := rule8[i]; !ok {
						rule8[i] = 0
					}
					rule8[i] += 1
					if rule8[i] > max {
						return "(42)" // 8 always returns 42
					}
				}
				if i == 11 {
					if _, ok := rule11[i]; !ok {
						rule11[i] = 0
					}
					rule11[i] += 1
					if rule11[i] > max {
						return "(42 31)"
					}
				}
				return rules[i]
			})
		}
		if done {
			break
		}
	}

	regexRules := map[int]*regexp.Regexp{}

	for k := range rules {
		// replacer will create Proper regex
		replacer := strings.NewReplacer("\"", "", " ", "")
		regexRules[k] = regexp.MustCompile("^" + replacer.Replace(rules[k]) + "$")
	}

	first := regexRules[0]

	count := 0
	for _, m := range messages {
		if first.MatchString(m) {
			count += 1
		}
	}
	return count

}

func parseInput(input []string) (rules map[int]string, messages []string) {

	rules = map[int]string{}
	parts := strings.Split(strings.Join(input, "\n"), "\n\n")
	for _, line := range strings.Split(parts[0], "\n") {
		k := adventofgo.AsInt(strings.Split(line, ":")[0])
		// wrap with () to help with generated regex
		v := "(" + strings.Split(line, ":")[1][1:] + ")"
		rules[k] = v
	}
	messages = strings.Split(parts[1], "\n")
	return

}
