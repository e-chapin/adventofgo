package main

import (
	"adventofgo"
	"fmt"
)

type seatCoord struct {
	row int
	col int
}

func main() {

	strlines := adventofgo.ReadFile("input.txt")
	seatMap := make(map[seatCoord]string, len(strlines))

	for ri, row := range strlines {
		for ci, column := range row {
			seat := string(column)
			seatMap[seatCoord{ri, ci}] = seat
		}
	}

	var seats int

	initialSeats := copySeatMap(seatMap)

	for {
		var changes int
		changes, seats, seatMap = checkAdjacentSeats(seatMap)
		if changes == 0 {
			break
		}
	}
	fmt.Println("Day 11 Part 1")
	fmt.Println(seats)

	seatMap = initialSeats
	for {
		var changes int
		changes, seats, seatMap = checkVisibleSeats(seatMap)
		if changes == 0 {
			break
		}
	}

	fmt.Println("Day 11 Part 2")
	fmt.Println(seats)

}

func copySeatMap(seatMap map[seatCoord]string) map[seatCoord]string {
	tempSeatMap := make(map[seatCoord]string)
	for i := range seatMap {
		tempSeatMap[i] = seatMap[i]
	}
	return tempSeatMap
}

func checkAdjacentSeats(seatMap map[seatCoord]string) (int, int, map[seatCoord]string){
	seats, changes := 0, 0

	tempSeatMap := copySeatMap(seatMap)

	for coordiate := range seatMap {
		seat := seatMap[coordiate]
		switch seat {
		case ".":
			// floor
			continue
		case "L":
			// empty seat
			if countAdjacent(coordiate, seatMap) == 0 {
				tempSeatMap[coordiate] = "#"
				changes += 1
				// newly occupied
				seats += 1
			}
		case "#":
			if countAdjacent(coordiate, seatMap) >= 4 {
				tempSeatMap[coordiate] = "L"
				changes += 1
			} else {
				// still occupied
				seats += 1
			}
		}
	}
	return changes, seats, tempSeatMap
}

func countAdjacent(coordinate seatCoord, grid map[seatCoord]string) int{

	count := 0
	checkSeatCount := 0

	for _, rdiff := range []int{-1, 0, 1} {
		for _, cdiff := range []int{-1, 0, 1} {

			checkSeatCount += 1

			r := coordinate.row+rdiff
			c := coordinate.col+cdiff

			if r < 0 || c < 0 || (cdiff == 0 && rdiff == 0){
				continue
			}

			gridVal := grid[seatCoord{r, c}]
			if gridVal == "#" {
				count += 1
			}
		}
	}
	return count

}

func checkVisibleSeats(seatMap map[seatCoord]string) (int, int, map[seatCoord]string){
	seats, changes := 0, 0

	tempSeatMap := copySeatMap(seatMap)

	for coordinate := range seatMap {

		seat := seatMap[coordinate]
		switch seat {
		case ".":
			// floor
			continue
		case "L":
			// empty seat
			if countVisibleSeats(coordinate, seatMap) == 0 {
				tempSeatMap[coordinate] = "#"
				changes += 1
				// newly occupied
				seats += 1
			}
		case "#":
			if countVisibleSeats(coordinate, seatMap) >= 5 {
				tempSeatMap[coordinate] = "L"
				changes += 1
			} else {
				// still occupied
				seats += 1
			}
		}
	}
	return changes, seats, tempSeatMap
}

func countVisibleSeats(coordinate seatCoord, grid map[seatCoord]string) int{

	count := 0
	for _, rdiff := range []int{-1, 0, 1} {
		for _, cdiff := range []int{-1, 0, 1} {

			if cdiff == 0 && rdiff == 0{
				continue
			}

			r := coordinate.row
			c := coordinate.col

			for {
				r = r+rdiff
				c = c+cdiff

				gridVal := grid[seatCoord{r, c}]

				if gridVal == "" || gridVal == "L" {
					break
				}

				if gridVal == "#" {
					count += 1
					break
				}
			}
		}
	}
	return count

}