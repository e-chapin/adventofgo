package main

import (
	"adventofgo"
	"fmt"
	"strconv"
)


func findSum(sum int, lines []int) bool {

	// probably a more clever way to do this.
	for i := 0; i < len(lines); i++ {
		first := lines[i]
		for j := 0; j < len(lines); j++ {
			second := lines[j]
			if sum == first+second {
				return true
			}
		}
	}
	return false

}

func main() {

	strlines := adventofgo.ReadFile("input.txt")
	lines := make([]int, len(strlines))

	for index, line := range strlines {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}
		lines[index] = n
	}

	fmt.Println("Part 1")

	preamble := 25

	for i := 0; i < len(lines); i++ {
		if i < preamble {
			continue
		}
		number := lines[i]
		if !findSum(number, lines[i-preamble:i]){
			fmt.Println(number)
			break
		}


	}
}
