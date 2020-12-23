package main

import (
	"adventofgo"
	"container/ring"
	"fmt"
)

func getStartingCups(input string) (cups *ring.Ring) {

	cups = ring.New(len(input))
	for _, v := range input {
		cups.Value = adventofgo.AsInt(string(v))
		cups = cups.Next()
	}
	return
}

func getMillionCups(input string) (cups *ring.Ring) {

	cups = ring.New(1000000)
	for _, v := range input {
		cups.Value = adventofgo.AsInt(string(v))
		cups = cups.Next()
	}
	for i := len(input)+1; i <= 1000000; i +=1 {
		cups.Value = i
		cups = cups.Next()
	}
	return
}

func playGame(cups *ring.Ring, rounds int) *ring.Ring {

	var max int
	l := cups.Len()

	// start by building a hash of cup values and pointers to a cup, so we don't have to search
	// sequentially every move. Also set max value.
	cupMap := make(map[int]*ring.Ring, l)
	for i:= 0; i < l; i += 1 {
		if cups.Value.(int) > max {
			max = cups.Value.(int)
		}
		cupMap[cups.Value.(int)] = cups
		cups = cups.Next()
	}

	for i := 0; i < rounds; i +=1 {
		// remove 3 to the right of current cup
		removedCups := cups.Unlink(3)

		destination := cups.Value.(int)
		for {
			destination -= 1
			if destination < 1 {
				destination = max
			}
			inRemovedCups := false
			removedCups.Do(func(v interface{}){
				if v.(int) == destination {
					inRemovedCups = true
				}
			})
			if !inRemovedCups {
				break
			}
		}

		// use the map to find the destination cup, and insert the removed next to it
		cupMap[destination].Link(removedCups)
		// move one cup right
		cups = cups.Next()
	}

	// return the spot in the ring of the first cup
	return cupMap[1]
}

func main() {

	input := adventofgo.GetFileString("input.txt")
	cups := getStartingCups(input)

	fmt.Println("Day 23 Part 1")
	cups = playGame(cups, 10)
	cups.Do(func(v interface{}){
		if v.(int) != 1 {
			fmt.Print(v.(int))
		}
	})
	fmt.Println()
	cups = getMillionCups(input)


	fmt.Println("Day 23 Part 2")
	cups = playGame(cups, 10000000)
	fmt.Println(cups.Next().Value.(int) * cups.Next().Next().Value.(int))
}
