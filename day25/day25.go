package main

import (
	"adventofgo"
	"fmt"
)

func main() {
	input := adventofgo.ReadFile("input.txt")
	fmt.Println("Day 25 Part 1")
	var card = adventofgo.AsInt(input[0])
	var door = adventofgo.AsInt(input[1])

	loopCard := 0
	for keyCard := 1; keyCard != card; loopCard += 1 {
		keyCard = (keyCard*7)%20201227
	}
	
	loopDoor := 0
	for keyDoor := 1; keyDoor != card; loopDoor++ {
		keyDoor = keyDoor * 7 % 20201227
	}

	encKey1 := 1
	for l := 0; l < loopCard; l++ {
		encKey1 = encKey1 * door % 20201227
	}
	encKey2 := 1
	for l := 0; l < loopDoor; l++ {
		encKey2 = encKey2 * door % 20201227
	}

	fmt.Println(encKey1)
	fmt.Println(encKey2)

}
