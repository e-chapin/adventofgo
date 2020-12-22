package main

import (
	"adventofgo"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func getDeck(input string) (deck []int) {

	for _, v := range strings.Split(input, "\n")[1:] {
		deck = append(deck, adventofgo.AsInt(v))
	}
	return
}

func getSubDeck(deck []int, card int) (sub []int) {
	for i := 0; i < card; i++ {
		sub = append(sub, deck[i])
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

func deckToHash(deck []int) string{
	strDeck := ""
	for i := range deck {
		strDeck += strconv.Itoa(deck[i])
	}
	hash := md5.Sum([]byte(strDeck))
	return hex.EncodeToString(hash[:])
}

func playGame(p1, p2 []int, game int) int {

	p1History := make(map[string]bool)
	p2History := make(map[string]bool)

	// main game loop for each game
	for len(p1) > 0 && len(p2) > 0 {
		var winner int
		p1DeckVal := deckToHash(p1)
		p2DeckVal := deckToHash(p2)

		if p1History[p1DeckVal] || p2History[p2DeckVal] {
			return 1
		} else {
			p1History[p1DeckVal] = true
			p2History[p2DeckVal] = true
		}

		c1, c2 := p1[0], p2[0]
		p1, p2 = p1[1:], p2[1:]


		// game -1 means this is part one, and we aren't playing Recursive Combat
		// If both player 1 and 2 have *equal or more* cards in their deck than the card they drew
		// then play a sub-game of Recursive Combat
		if len(p1) >= c1 && len(p2) >= c2 && game > 0 {

			s1 := getSubDeck(p1, c1)
			s2 := getSubDeck(p2, c2)
			winner = playGame(s1, s2, game+1)

		// otherwise just play a round of normal Combat
		} else {
			if c1 > c2 {
				winner = 1
			} else {
				winner = 0
			}
		}

		if winner == 1 {
			// Player 1 wins. add cards to bottom of player 1 deck, keeping c1 and then adding c2.
			p1 = append(p1, c1, c2)
		} else {
			// player 2 wins, do opposite.
			p2 = append(p2, c2, c1)
		}
	}

	if len(p1) > 0 {
		if game <= 1 {
			fmt.Println(CalculateScore(p1))
		}
		return 1
	} else {
		if game <= 1 {
			fmt.Println(CalculateScore(p2))
		}
		return 0
	}
}

func main() {

	input := adventofgo.GetFileString("input.txt")

	players := strings.Split(input, "\n\n")
	playerOne := getDeck(players[0])
	playerTwo := getDeck(players[1])

	fmt.Println("Day 22 Part 1")
	_ = playGame(playerOne, playerTwo, -1)
	fmt.Println("Day 22 Part 2")
	_ = playGame(playerOne, playerTwo, 1)


}
