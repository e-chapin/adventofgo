package main

import (
	"adventofgo"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type bits struct {
	bitmap map[int]int
}


//type address struct {
//	value []bit
//	index int
//}

func main() {


	address := map[int32]bits{}

	for i:= 0; i < 36; i++ {
		var b = bits{bitmap: make(map[int]int)}
		for j:= 1; j < 34359738368; j*=2 {
			b.bitmap[j] = 0
		}
		address[int32(i)] = b
	}


	strlines := adventofgo.ReadFile("input.txt")

	var mask string
	mask = ""
	for _, val := range strlines {
		splLine := strings.Split(val, " = ")
		if splLine[0] == "mask" {
			fmt.Println("old mask", mask)
			mask = splLine[1]
			fmt.Println("new mask", mask)
			continue
		}
		// mem[8] = 11
		value, _ := strconv.Atoi(splLine[1])
		maxIndex := len(splLine[0])-1
		stradr := string(splLine[0][4:maxIndex])
		adr, _ := strconv.Atoi(stradr)

		fmt.Println(value)
		fmt.Println(adr)

		newValueBits := extrapValue(value)
		maskedBits := applyValueWithMask(newValueBits, mask)
		address[int32(adr)] = maskedBits
	}
	fmt.Println("Day 14 Part 1")
	fmt.Println(memorySum(address))

}

func extrapValue(value int) bits {

	var b = bits{bitmap: make(map[int]int)}

	for j:= 34359738368; j >= 1; j=j/2 {
		b.bitmap[j] = 0
		if j <= value {
			b.bitmap[j] = 1
			value -= j
		}
	}
	return b
}

func applyValueWithMask(value bits, mask string) bits{

	var b = bits{bitmap: make(map[int]int)}

	//apply Mask to value
	mask = adventofgo.ReverseString(mask)
	for j := 1; j < 34359738368; j*=2 {
		bitIndex := int(math.Log2(float64(j)))

		if string(mask[bitIndex]) == "X" {
			b.bitmap[j] = value.bitmap[j]
			continue
		} else {
			newValStr := string(mask[bitIndex])
			newVal, _ := strconv.Atoi(newValStr)
			b.bitmap[j] = newVal
		}
	}

	return b
}

func memorySum(address map[int32]bits) int {

	var sum int
	sum = 0
	for i:= 0; i < 36; i++ {
		for j:= 1; j < 34359738368; j*=2 {
			if address[int32(i)].bitmap[j] == 1 {
				sum += j
			}
		}
	}
	return sum
}