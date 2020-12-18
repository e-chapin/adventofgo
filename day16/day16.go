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
	maxtwo int
}

type ticket struct {

	values []int

}


func main() {

	var tix []ticket

	myTicket, validTickets := findInvalidTickets()

	for _, row := range validTickets {
		fmt.Println(row)
		var t ticket
		err := json.Unmarshal([]byte("["+row+"]"), &t.values)
		if err != nil {
			log.Fatal(err)
		}
		tix = append(tix, t)
		fmt.Printf("%v", t)

	}


	fmt.Println(myTicket)
	//fmt.Println(len(validTickets))

}

func findInvalidTickets() ([]int, []string) {
	var rules = map[int]bool{}
	var myTicket = []int{}
	var otherTickets = []string{}

	var ticketFields []field

	strlines := adventofgo.ReadFile("input.txt")
	var index = 0
	for {
		val := strlines[index]
		index += 1
		if val == "" {
			break
		}

		r := regexp.MustCompile(`(?P<label>[a-z]+): (?P<minOne>\d+)-(?P<maxOne>\d+) or (?P<minTwo>\d+)-(?P<maxTwo>\d+)`)
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
			maxtwo: iMaxTwo,
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

	}

	otherTickets = append(otherTickets, strlines[index+1:]...)

	var totalInvalid = 0
	var validTickets = []string{}
	for _, ticket := range otherTickets {
		var valid = true
		fields := strings.Split(ticket, ",")
		for _, field := range fields {
			ifield, _ := strconv.Atoi(field)
			if val, ok := rules[ifield]; !ok || !val {
				valid = false
				totalInvalid += ifield
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}

	}

	fmt.Println("Day 16 Part 1")
	fmt.Println(totalInvalid)
	return myTicket, validTickets
}
