package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(0)
	}

	lines := strings.Split(string(content), "\n")

	valid := 0
	valid_pt2 := 0

	for _, line := range lines{

		splitline := strings.Split(line, " ")

		minmax := strings.Split(splitline[0], "-")

		min, _ := strconv.Atoi(minmax[0])
		max, _ := strconv.Atoi(minmax[1])

		letter := strings.Split(splitline[1], ":")[0]

		pw := splitline[2]

		count := strings.Count(pw, letter)

		if count >= min && count <= max {
			valid += 1
		}

		char_count := 0

		if string(pw[min-1]) == letter {
			char_count++
		}
		if string(pw[max-1]) == letter{
			char_count++
		}
		if char_count == 1 {
			valid_pt2 += 1
		}
	}

	fmt.Println(valid)
	fmt.Println(valid_pt2)

}
