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

	strlines := adventofgo.ReadFile("input.txt")
	var mask string
	for _, val := range strlines {
		splLine := strings.Split(val, " = ")
		if splLine[0] == "mask" {
			mask = splLine[1]
			continue
		}

		// this is sad and can be improved
		value, _ := strconv.ParseInt(splLine[1], 10, 64)

		adr, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(splLine[0], "mem["), "]"))
		memory[adr] = maskValue(value, mask)

	}
	fmt.Println("Day 14 Part 1")
	answer := 0
	for _, v := range memory {
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
		fmt.Println(maskChar, value, currentBinaryValue)
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