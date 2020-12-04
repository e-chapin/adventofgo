package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(0)
	}

	lines := strings.Split(string(content), "\n")
	max := len(lines)

	for i := 0; i < max; i++ {
		a, err := strconv.Atoi(lines[i])
		if err != nil {
			fmt.Println(err)
		}
		for j := i+1; j < max ; j++ {
			b, err := strconv.Atoi(lines[j])
			if err != nil {
				fmt.Println(err)
			}
			if a + b == 2020 {
				fmt.Println("part 1")
				fmt.Println(a*b)
				break
			}
			for z := j + 1; z < max; z++ {
				c, err := strconv.Atoi(lines[z])
				if err != nil {
					fmt.Println(err)
				}
				if a+b+c == 2020 {
					fmt.Println("part 2")
					fmt.Println(a * b * c)
					break
				}
			}
		}
	}

}
