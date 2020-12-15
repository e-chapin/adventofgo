package main

import (
	"adventofgo"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	memory := make(map[int]int64)
	memory2 := make(map[int]int64)

	strlines := adventofgo.ReadFile("input.txt")
	var mask string
	for _, val := range strlines {
		splLine := strings.Split(val, " = ")
		if splLine[0] == "mask" {
			mask = splLine[1]
			continue
		}

		value, _ := strconv.ParseInt(splLine[1], 10, 64)

		adr, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(splLine[0], "mem["), "]"))
		// Part 1 is a single masking of the value
		memory[adr] = maskValue(value, mask)

		// Part 2 need to convert its mask to a list of part 1 masks
		masks := convertMask("", mask)
		for _, m := range masks {
			// and then apply each to the memory address, and put value at that address.
			maskedAddress := int(maskValue(int64(adr), m))
			memory2[maskedAddress] = value
		}

	}
	fmt.Println("Day 14 Part 1")
	answer := 0
	for _, v := range memory {
		answer += int(v)
	}
	fmt.Println(answer)

	fmt.Println("Day 14 Part 2")
	answer = 0
	for _, v := range memory2 {
		answer += int(v)
	}
	fmt.Println(answer)

}

func maskValue(value int64, mask string) int64 {

	var maskedVal int64 = 0
	var currentBinaryValue int64
	for index, v := range adventofgo.ReverseString(mask){
		maskChar := string(v)
		currentBinaryValue = int64(math.Pow(2, float64(index)))
		if maskChar == "0" {
			// set as 0, do nothing
			continue
		} else if maskChar == "X" {
			maskedVal += value & currentBinaryValue
		} else if maskChar == "1" {
			//set as 1
			maskedVal += currentBinaryValue
		}
	}
	return maskedVal
}

// convert part 2 mask to list of part 1 mask
func convertMask(newMask, startMask string) []string{

	if len(startMask) == 0 {
		return []string{newMask}
	}

	switch startMask[0]{
	case '0':
		// 0 is the new X
		return convertMask(newMask+"X", startMask[1:])

	case '1':
		// 1 is the same
		return convertMask(newMask+"1", startMask[1:])

	case 'X':
		// X needs to be both 1 and 0 in the final list
		maskWithZero := convertMask(newMask+"0", startMask[1:])
		maskWithOne := convertMask(newMask+"1", startMask[1:])
		return append(maskWithZero, maskWithOne...)

	default:
		// bad val
		return []string{""}

	}

}