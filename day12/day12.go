package main

import (
	"adventofgo"
	"fmt"
	"strconv"
)



func main() {

	strlines := adventofgo.ReadFile("input.txt")

	east := 0
	north := 0

	wpNorth := 1
	wpEast := 10

	var orientation string
	//orientation_key := 3
	direction := []string{"N", "E", "S", "W"}
	orientation = direction[1]

	part1(strlines, orientation, north, east, direction)

	part2(strlines, east, wpEast, north, wpNorth)


}

func part1(strlines []string, orientation string, north int, east int, direction []string) {
	for _, line := range strlines {

		action := line[:1]
		value, _ := strconv.Atoi(line[1:])

		switch action {
		case "F":
			switch orientation {
			case "N":
				north += value
			case "E":
				east += value
			case "S":
				north -= value
			case "W":
				east -= value
			}

		case "N":
			north += value
		case "S":
			north -= value
		case "E":
			east += value
		case "W":
			east -= value

		case "R", "L":

			// I think this could be simpler?
			index := adventofgo.IndexOf(orientation, direction)
			offset := value / 90

			if action == "L" {
				offset = offset * -1
			}

			newIndex := (index + offset) % 4
			for {
				if newIndex >= 0 {
					break
				}
				newIndex += 4
				if newIndex == 4 {
					newIndex = 0
				}
			}
			orientation = direction[newIndex]

		}
	}

	fmt.Println("Day 12 Part 1")
	fmt.Println(adventofgo.Abs(north) + adventofgo.Abs(east))
}

func part2(strlines []string, east int, wpEast int, north int, wpNorth int){
	for _, line := range strlines {
		action := line[:1]
		value, _ := strconv.Atoi(line[1:])

		switch action {
		case "F":
			east += value * wpEast
			north += value * wpNorth

		case "N":
			wpNorth += value
		case "S":
			wpNorth -= value
		case "E":
			wpEast += value
		case "W":
			wpEast -= value

		case "R":
			turns := value / 90
			for i := 0; i < turns; i++ {
				wpNorth, wpEast = wpEast*-1, wpNorth
			}

		case "L":
			turns := value / 90
			for i := 0; i < turns; i++ {
				wpNorth, wpEast = wpEast, wpNorth*-1
			}
		}
	}
	fmt.Println("Day 12 Part 2")
	fmt.Println(adventofgo.Abs(north) + adventofgo.Abs(east))
}
