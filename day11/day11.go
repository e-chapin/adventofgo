package main

import (
	"adventofgo"
	"fmt"
)

func main() {

	fmt.Println("Day 11 Part 1")

	strlines := adventofgo.ReadFile("input.txt")
	var seatMap [][]string
	seatMap = make([][]string, len(strlines))

	for row, line := range strlines {
		seatMap[row] = make([]string, len(line))
		for column, s := range line {
			seat := string(s)
			seatMap[row][column] = seat
		}
	}

	var seats int
	for {
		var changes int
		changes, seats, seatMap = checkSeats(seatMap)
		if changes == 0 {
			break
		}
	}
	fmt.Println(seats)

}


func checkSeats(seatMap [][]string) (int, int, [][]string){
	seats, changes := 0, 0

	tempSeatMap := copySeatMap(seatMap)

	for row, line := range seatMap {
		for column, seat := range line {

			switch seat {
			case ".":
				// floor
				continue
			case "L":
				// empty seat
				if countAdjacent(row, column, seatMap) == 0 {
					tempSeatMap[row][column] = "#"
					changes += 1
					// newly occupied
					seats += 1
				}
			case "#":
				if countAdjacent(row, column, seatMap) >= 4 {
					tempSeatMap[row][column] = "L"
					changes += 1
				} else {
					// still occupied
					seats += 1
				}
			}
		}
	}
	return changes, seats, tempSeatMap
}

func copySeatMap(seatMap [][]string) [][]string {
	tempSeatMap := make([][]string, len(seatMap))
	for i := range seatMap {
		tempSeatMap[i] = make([]string, len(seatMap[i]))
		copy(tempSeatMap[i], seatMap[i])
	}
	return tempSeatMap
}

func countAdjacent(row int, column int, grid [][]string) int{

	count := 0
	checkSeatCount := 0
	for _, rdiff := range []int{-1, 0, 1} {
		for _, cdiff := range []int{-1, 0, 1} {

			checkSeatCount += 1

			r := row+rdiff
			c := column+cdiff

			if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || (cdiff == 0 && rdiff == 0){
				continue
			}

			gridVal := grid[r][c]
			if gridVal == "#" {
				count += 1
			}
		}
	}
	return count

}