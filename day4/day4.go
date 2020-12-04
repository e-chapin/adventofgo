package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkPassportPartOne(passport map[string]string) bool {
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, key := range keys {
		if _, ok := passport[key]; !ok {
			return false
		}
	}
	return true
}


func checkPassportPartTwo(passport map[string]string) bool {

	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, key := range keys {
		if _, ok := passport[key]; !ok {
			return false
		}

		switch key {
		case "byr":
			year, err := strconv.Atoi(passport[key])
			if err != nil || year < 1920 || year > 2002 {
				return false
			}

		case "iyr":
			year, err := strconv.Atoi(passport[key])
			if err != nil || year < 2010 || year > 2020 {
				return false
			}

		case "eyr":
			year, err := strconv.Atoi(passport[key])
			if err != nil || year < 2020 || year > 2030 {
				return false
			}

		case "hgt":
			height := passport[key]
			number, err := strconv.Atoi(height[:len(height)-2])
			if err != nil{
				return false
			}
			unit := height[len(height)-2:]

			switch unit{
			case "cm":
				if number < 150 || number > 193{
					return false
				}
			case "in":
				if number < 59 || number > 76{
					return false
				}
			default:
				return false

			}

		case "hcl":
			color := passport[key]
			matched, err := regexp.MatchString("#[a-f0-9]{6}", color)
			if err != nil || !matched || len(color) != 7{
				return false
			}

		case "ecl":
			color := passport[key]
			matched, err := regexp.MatchString("amb|blu|brn|gry|grn|hzl|oth", color)
			if err != nil || !matched || len(color) != 3{
				return false
			}

		case "pid":
			id := passport[key]
			matched, err := regexp.MatchString("[0-9]{9}", id)
			if err != nil || !matched || len(id) != 9{
				return false
			}

		}

	}
	return true
}


func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(0)
	}

	lines := strings.Split(string(content), "\n")

	currentPassport := map[string]string{}
	validCountPartOne := 0
	validCountPartTwo := 0

	for index, line := range lines {

		if line == ""{
			// check for completed passport
			if checkPassportPartOne(currentPassport){
				validCountPartOne += 1
			}
			if checkPassportPartTwo(currentPassport){
				validCountPartTwo += 1
			}
			currentPassport = make(map[string]string)
			continue
		}

		fields := strings.Split(line, " ")

		for _, field := range fields {

			split := strings.Split(field, ":")

			key := split[0]
			value := split[1]

			currentPassport[key] = value

		}

		// if this is the last line, check the last passport
		if index == len(lines)-1 {
			if checkPassportPartOne(currentPassport){
				validCountPartOne += 1
			}
			if checkPassportPartTwo(currentPassport){
				validCountPartTwo += 1
			}
			break
		}

	}

	fmt.Println("Part 1")
	fmt.Println(validCountPartOne)
	fmt.Println("Part 2")
	fmt.Println(validCountPartTwo)

}
