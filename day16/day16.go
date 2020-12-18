package main

import (
	"adventofgo"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type field struct {
	label string
	minOne int
	maxOne int
	minTwo int
	maxTwo int
}

type ticket struct {

	values []int

}


var validTickets []ticket
var rules = map[int]bool{}
var myTicket []int
var otherTickets []string
var ticketFields []field
//var validTickets []string
var invalidSum int

var fieldPos = make(map[int]int)

var myTicketFielPos = make(map[string]int)

func main() {

	populateDataStructures()

	invalidSum = findInvalidTickets()
	fmt.Println("Day 16 part 1")
	fmt.Println(invalidSum)

	fmt.Println(myTicket)

	findTicketLabels()
}


func allValuesMatch(input []int, f field) bool {

	allMatch := true
	for _, i := range input {
		if !((i <= f.maxOne && i >= f.minOne) || ( i <= f.maxTwo && i >= f.minTwo)) {
			allMatch = false
			break
		}
	}
	return allMatch

}

func findTicketLabels() {

		var foundIndex []int

		for index := 0; index < len(ticketFields); index++ {
			var vals []int
			// get all ticket vals at this index
			for _, t := range validTickets {
				vals = append(vals, t.values[index])
			}

			for ticketIndex, f := range ticketFields {

				//if adventofgo.IndexOfInt(ticketIndex, foundIndex) >= 0 {
				//	continue
				//}

				if allValuesMatch(vals, f) {
					fieldPos[ticketIndex] = index
					myTicketFielPos[f.label] = myTicket[index]
					fmt.Println("Found match", f.label, myTicket[index])
					foundIndex = append(foundIndex, ticketIndex)
					break
				}
			}

			fmt.Println(vals)
			fmt.Println(vals)

		}
		fmt.Println("stop")
		fmt.Println(myTicketFielPos)

		product := 1
		for k, v := range myTicketFielPos {
			if strings.HasPrefix(k, "departure"){
				product *= v
			}
		}
		fmt.Println("Day 16 part 2")
		fmt.Println(product)

}

func findInvalidTickets() int {

	var totalInvalid = 0
	for _, tck := range otherTickets {
		var valid = true
		fields := strings.Split(tck, ",")
		for _, field := range fields {
			ifield, _ := strconv.Atoi(field)
			if val, ok := rules[ifield]; !ok || !val {
				valid = false
				totalInvalid += ifield
			}
		}
		if valid {
			// populate the last data struct
			var t ticket
			err := json.Unmarshal([]byte("["+tck+"]"), &t.values)
			if err != nil {
				log.Fatal(err)
			}
			validTickets = append(validTickets, t)
		}
	}

	fmt.Println("Day 16 Part 1")
	fmt.Println(totalInvalid)
	return totalInvalid
}

func populateDataStructures() {

	strlines := adventofgo.ReadFile("input.txt")
	var index = 0
	for {
		val := strlines[index]
		index += 1
		if val == "" {
			break
		}

		r := regexp.MustCompile(`(?P<label>[a-z ]+): (?P<minOne>\d+)-(?P<maxOne>\d+) or (?P<minTwo>\d+)-(?P<maxTwo>\d+)`)
		match := r.FindStringSubmatch(val)
		result := make(map[string]string)
		for i, name := range r.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		iMinOne, _ := strconv.Atoi(result["minOne"])
		iMaxOne, _ := strconv.Atoi(result["maxOne"])
		iMinTwo, _ := strconv.Atoi(result["minTwo"])
		iMaxTwo, _ := strconv.Atoi(result["maxTwo"])

		f := field{
			label:  result["label"],
			minOne: iMinOne,
			maxOne: iMaxOne,
			minTwo: iMinTwo,
			maxTwo: iMaxTwo,
		}

		ticketFields = append(ticketFields, f)

		for i := iMinOne; i <= iMaxOne; i++ {
			rules[i] = true
		}
		for i := iMinTwo; i <= iMaxTwo; i++ {
			rules[i] = true
		}

	}

	for {
		val := strlines[index]
		index += 1
		if val == "" {
			break
		}
		if val == "your ticket:" {
			continue
		}
		for _, v := range strings.Split(val, ",") {
			v, _ := strconv.Atoi(v)
			myTicket = append(myTicket, v)
		}
		validTickets = append(validTickets, ticket{myTicket})

	}

	otherTickets = append(otherTickets, strlines[index+1:]...)
}
