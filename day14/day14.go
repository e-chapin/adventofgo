package main

import (
	"adventofgo"
	"fmt"
	"strconv"
	"strings"
)

func main() {


	memory := make(map[int]int64)

	strlines := adventofgo.ReadFile("input.txt")

	var andMask int64 = 0
	var orMask int64 = 0
	for _, val := range strlines {
		splLine := strings.Split(val, " = ")
		if splLine[0] == "mask" {
			// first mask, both value and mask need to have a 1, and each "X" is converted to a 1.
			// Guarantees that any matching 1's from value and map are retained by X's in map,
			// but that zeroes in map transfer to value and replace 1s.
			andMask, _ = strconv.ParseInt(strings.ReplaceAll(splLine[1], "X", "1"), 2, 0)

			// second mask, only one of them need to have a 1, and each "X" is converted to a 0.
			// This will guarantee all hard 1s in Mask will overwrite into value if that spot is a 0 in value.
			orMask, _ = strconv.ParseInt(strings.ReplaceAll(splLine[1], "X", "0"), 2, 0)
			continue
		}


		// this is sad and can be improved
		intval, _ := strconv.Atoi(splLine[1])
		i64Val := int64(intval)
		maxIndex := len(splLine[0])-1
		stradr := string(splLine[0][4:maxIndex])
		adr, _ := strconv.Atoi(stradr)

		i64Val &= andMask
		i64Val |= orMask
		memory[adr] = i64Val

	}
	fmt.Println("Day 14 Part 1")
	answer := 0
	for _, v := range memory {
		answer += int(v)
	}
	fmt.Println(answer)
}