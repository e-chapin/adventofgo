package main

import (
	"adventofgo"
	"fmt"
	"github.com/deanveloper/modmath/v1/bigmod"
	"math/big"
	"strconv"
	"strings"
)


func main() {


	strlines := adventofgo.ReadFile("input.txt")

	departTime, _ := strconv.Atoi(strlines[0])
	busSchedule := strlines[1]

	busId := 999999
	minBusDepartTime := 999999999

	for _, val := range strings.Split(busSchedule, ",") {

		routeTime, _ := strconv.Atoi(val)

		if val == "x" {
			continue
		}

		busDepartTime := departTime-(departTime%routeTime)+routeTime

		if busDepartTime < minBusDepartTime{
			minBusDepartTime = busDepartTime
			busId = routeTime
		}
	}

	fmt.Println("Day 13 Part 1")
	fmt.Println(busId*(minBusDepartTime - departTime))

	// Part 2?

	var busses []bigmod.CrtEntry

	for index, val := range strings.Split(busSchedule, ",") {
		if val == "x" {
			continue
		}
		busId, _ := strconv.Atoi(val)
		busEntry := bigmod.CrtEntry{A: big.NewInt(int64(busId-index)), N: big.NewInt(int64(busId))}
		busses = append(busses, busEntry)

		fmt.Println(busId-index, busId)

	}
	fmt.Println("Day 13 part 2")
	fmt.Println(bigmod.SolveCrtMany(busses))

}
