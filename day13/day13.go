package main

import (
	"adventofgo"
	"fmt"
	"strconv"
	"strings"
)

func main() {


	strlines := adventofgo.ReadFile("input.txt")

	departTime, _ := strconv.Atoi(strlines[0])
	busSchedule := strlines[1]

	fmt.Println(departTime)
	fmt.Println(busSchedule)

	busId := 999999
	minBusDepartTime := 999999999

	for index, val := range strings.Split(busSchedule, ",") {

		routeTime, _ := strconv.Atoi(val)

		if val == "x" {
			continue
		}

		busDepartTime := departTime-(departTime%routeTime)+routeTime

		if busDepartTime < minBusDepartTime{
			minBusDepartTime = busDepartTime
			busId = routeTime
		}


		fmt.Println(index)
		fmt.Println(val)
	}

	fmt.Println(busId)
	fmt.Println(minBusDepartTime - departTime)
	fmt.Println(busId*(minBusDepartTime - departTime))

}
