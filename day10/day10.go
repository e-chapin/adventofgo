package main

import (
	"adventofgo"
	"fmt"
	"sort"
	"strconv"
)

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

	sort.Ints(lines)
	lines = append([]int{0}, lines...)
	lines = append([]int{lines[len(lines)-1]+3}, lines...)
	sort.Ints(lines)

	differences := make(map[int]int)


	for index, line := range lines {
		if index == 0 {
			continue
		}
		differences[line - lines[index-1]]++
	}

	fmt.Println(differences[1]*differences[3])


}
