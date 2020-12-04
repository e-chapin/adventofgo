package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(0)
	}

	lines := strings.Split(string(content), "\n")
	maxSize := len(lines[0])

	slopes := [][]int{
		{1, 1}, {3, 1}, {5, 1},{7, 1}, {1, 2},
	}

	mult_total := 1

	for _, slope := range slopes {

		column := 0
		right := slope[0]
		down := slope[1]

		trees := 0

		for i, line := range lines {

			if i % down != 0 {
				continue
			}

			if i != 0 {
				column = (column + right) % maxSize
			}

			if line[column] == '#' {
				trees += 1
			}
			}

		if right == 3 {
			fmt.Println("Part 1")
			fmt.Println(trees)
		}

		mult_total = mult_total*trees

	}

	fmt.Println("Part 2")
	fmt.Println(mult_total)

}
