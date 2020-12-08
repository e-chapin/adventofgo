package main

import (
	"adventofgo"
	"fmt"
	"strconv"
	"strings"
)


var acc = 0


func main() {

	lines := adventofgo.ReadFile("input.txt")

	instRan := make(map[int]bool)

	for i:= 0; i < len(lines); i++ {


		if instRan[i] == true {
			fmt.Println("Part 1")
			fmt.Println(acc)
			break

		}

		line := lines[i]

		opp := strings.Split(line, " ")[0]
		
		switch opp {
		case "acc":
			arg := strings.Split(line, " ")[1]
			accumulator, _ := strconv.Atoi(arg)
			acc += accumulator
		case "jmp":
			arg := strings.Split(line, " ")[1]
			jmpVal, _ := strconv.Atoi(arg)
			i += jmpVal-1
		case "nop":
			continue
		}

		instRan[i] = true

	}

}
