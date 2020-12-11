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
		//fmt.Println("Row", row)
		for column, s := range line {
			seat := string(s)
			//fmt.Println("column", column)
			//fmt.Println("seat", seat)
			seatMap[row][column] = seat

		}
		//fmt.Println(index)
		//fmt.Println(line)
	}

	count := 0
	for {
		var changes int
		//tempSeatMap := copySeatMap(seatMap)
		changes, seatMap = checkSeats(seatMap)
		fmt.Println(seatMap)
		if changes == 0 {
			break
		}
		count += 1
		//seatMap = copySeatMap(tempSeatMap)
	}
	fmt.Println(countSeats(seatMap))

}

func countSeats(seatMap [][]string) int{
	count := 0
	for _, line := range seatMap {
		for _, seat := range line {
			if seat == "#" {
				count += 1
			}
		}
	}
	return count
}

func checkSeats(seatMap [][]string) (int, [][]string){
	changes := 0

	tempSeatMap := copySeatMap(seatMap)

	for row, line := range seatMap {
		for column, seat := range line {

			switch seat {
			case ".":
				// floor
				continue
			case "L":
				// empty seat
				if countAdjecent(row, column, seatMap) == 0 {
					tempSeatMap[row][column] = "#"
					changes += 1
				}
			case "#":
				if countAdjecent(row, column, seatMap) >= 4 {
					tempSeatMap[row][column] = "L"
					changes += 1

				}
			}
		}
	}
	return changes, tempSeatMap
}

func copySeatMap(seatMap [][]string) [][]string {
	tempSeatMap := make([][]string, len(seatMap))
	for i := range seatMap {
		tempSeatMap[i] = make([]string, len(seatMap[i]))
		copy(tempSeatMap[i], seatMap[i])
	}
	return tempSeatMap
}

func countAdjecent(row int, column int, grid [][]string) int{

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