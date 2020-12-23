package main

import (
	"adventofgo"
	"fmt"
)

var max, min int

// LLNode is a doubly linked list node
type LL struct {
	val         int
	left, right *LL
}

func getStartingCups(input string) (cups []int) {

	for i := range input {
		val := adventofgo.AsInt(string(input[i]))
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
		cups = append(cups, val)
	}
	return
}

// 0 <= index <= len(a)
func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func doMove(cups []int, currentIndex int) ([]int, int){

	var removedCups []int

	currentCup := cups[currentIndex%len(cups)]
	fmt.Println("current cup", currentCup)

	for i := 1; i <= 3; i += 1 {
		index := (currentIndex+i)%len(cups)
		removedCups = append(removedCups, cups[index])
	}

	for _, v := range removedCups {
		i := adventofgo.IndexOfInt(v, cups)
		cups = append(cups[:i], cups[(i+1):]...)
	}

	var destinationIndex int
	destinationCup := currentCup
	for {
		destinationCup -= 1
		if destinationCup < min {
			destinationCup = max
		}
		destinationIndex = adventofgo.IndexOfInt(destinationCup, cups)
		if destinationIndex >= 0 {
			break
		}
	}

	fmt.Println("Destination cup", cups[destinationIndex])


	var tmpCups []int
	tmpCups = append(tmpCups, cups[:destinationIndex+1]...)
	tmpCups = append(tmpCups, removedCups...)
	tmpCups = append(tmpCups, cups[destinationIndex+1:]...)
	cups = tmpCups


	return cups, adventofgo.IndexOfInt(currentCup, cups)+1

}

func main() {

	input := adventofgo.GetFileString("input.txt")
	cups := getStartingCups(input)

	cupIndex := 0
	for i := 0; i < 100; i +=1 {

		fmt.Println("-- move", i+1, "--")
		fmt.Println(cups)
		cups, cupIndex = doMove(cups, cupIndex)
		fmt.Println()

	}
	fmt.Println("Day 23 Part 1")
	indexOne := adventofgo.IndexOfInt(1, cups)
	for i := indexOne+1; i < len(cups); i +=1 {
		fmt.Print(cups[i])
	}
	for i := 0; i < indexOne; i +=1 {
		fmt.Print(cups[i])
	}
	//fmt.Println(cups)


}
