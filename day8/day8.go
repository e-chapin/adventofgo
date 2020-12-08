package main

import (
	"adventofgo"
	"fmt"
	"strconv"
	"strings"
)

var acc = 0

func runLoop(lines []string) bool{
	instRan := make(map[int]bool)

	for i:= 0; i < len(lines); i++ {

		if instRan[i] == true {
			return false
		}
		instRan[i] = true

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

	}
	// finished commands without repeating, no infinite loop
	return true

}


func main() {

	lines := adventofgo.ReadFile("input.txt")

	fmt.Println("Part 1")
	runLoop(lines)
	fmt.Println(acc)

	//reset for part 2
	acc = 0

	for i := 0; i < len(lines); i++ {

		line := lines[i]
		opp := strings.Split(line, " ")[0]

		switch opp {
		case "nop":
			arg := strings.Split(line, " ")[1]
			line = "jmp " + arg
		case "jmp":
			arg := strings.Split(line, " ")[1]
			line = "nop " + arg
		}

		newLines := make([]string, len(lines))
		copy(newLines, lines)
		newLines[i] = line

		noLoop := runLoop(newLines); if noLoop {
			fmt.Println("Part 2")
			fmt.Println(acc)
		}else{
			acc = 0
		}
	}
}
