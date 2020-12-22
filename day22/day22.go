package main

import (
	"adventofgo"
	"fmt"
	"strings"
)

func getDeck(input string) (deck []int) {

	for _, v := range strings.Split(input, "\n")[1:] {
		deck = append(deck, adventofgo.AsInt(v))
	}
	return
}

func CalculateScore(deck []int) (score int){
	for i, v := range deck {
		factor := len(deck) - i
		score += v*factor
	}
	return
}

func playGame(p1, p2 []int) {

	for len(p1) > 0 && len(p2) > 0 {

		c1, c2 := p1[0], p2[0]
		p1 = p1[1:]
		p2 = p2[1:]

		if c1 > c2{
			// Player 1 wins. add cards to bottom of player 1 deck, keeping c1 and then adding c2.
			p1 = append(p1, c1)
			p1 = append(p1, c2)
		} else {
			// player 2 wins, do opposite.
			p2 = append(p2, c2)
			p2 = append(p2, c1)
		}
	}

	if len(p1) > 0 {
		fmt.Println(CalculateScore(p1))
	} else {
		fmt.Println(CalculateScore(p2))
	}

}

func main() {

	input := adventofgo.GetFileString("input.txt")

	players := strings.Split(input, "\n\n")
	playerOne := getDeck(players[0])
	playerTwo := getDeck(players[1])

	fmt.Println("Day 22 Part 1")
	playGame(playerOne, playerTwo)


}
