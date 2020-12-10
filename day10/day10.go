package main

import (
	"adventofgo"
	"fmt"
	"sort"
	"strconv"
)


var adapterSum = 0

//var max = -1

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
	max := lines[len(lines)-1] + 3
	lines = append([]int{0, max}, lines...)
	sort.Ints(lines)

	differences := make(map[int]int)

	for index, line := range lines {
		if index == 0 {
			continue
		}
		differences[line-lines[index-1]]++
	}
	fmt.Println("Part 1")
	fmt.Println(differences[1] * differences[3])

	fmt.Println("Part 2")
	fmt.Println(findGaps(lines))
}


// for every gap of between two elements, the only option is to use the +3 adapter, as +1 and +2 are guaranteed
// to not be valid adapters. Splice the big list until sub-lists for each of these, count individually and multiply.
func findGaps(adapterList []int) int {
	count := 1
	subIndex := 0

	for index, value := range adapterList {
		if index == len(adapterList)-1 {
			continue
		}
		nextIndex := index+1
		next := adapterList[nextIndex]
		if next - value == 3 {

			adapterSubset := adapterList[subIndex:nextIndex]
			count = count * countAdapters(adapterSubset[0], adapterSubset)
			subIndex = nextIndex
		}
	}
	return count
}


func countAdapters(input int, adapterList []int) int{

	localSum := 0
	max := adapterList[len(adapterList)-1]

	if input == max {
		adapterSum += 1
		return 1
	} else if !adventofgo.Contains(input, adapterList) || input > max {
		return 0
	} else {
			localSum += countAdapters(input+1, adapterList)
			localSum += countAdapters(input+2, adapterList)
			localSum += countAdapters(input+3, adapterList)
			return localSum
	}
}

