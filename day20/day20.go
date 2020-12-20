package main

import (
	"adventofgo"
	"fmt"
	"strings"
)

type  tile struct {
	// a tile has 8 sides, each side read left to right (forward) and then right to left (backward)
	number int
	tf int64
	tb int64
	rf int64
	rb int64
	bf int64
	bb int64
	lf int64
	lb int64
}

func main() {

	input := adventofgo.ReadFile("input.txt")
	countCorners(input)
}

func calculateTile(input string) tile {
	r := strings.NewReplacer("#", "1", ".", "0")
	input = r.Replace(input)

	lines := adventofgo.RemoveEmpty(strings.Split(input, "\n"))
	name := strings.TrimSuffix(strings.Split(lines[0], " ")[1], ":")

	top := lines[1]
	topBack := adventofgo.ReverseString(top)
	bottom := lines[len(lines)-1]
	bottomBack := adventofgo.ReverseString(bottom)
	left := ""
	right := ""

	for i:= 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		left += string(line[0])
		right += string(line[len(line)-1])
	}

	rightBack := adventofgo.ReverseString(right)
	leftBack := adventofgo.ReverseString(left)

	t := tile{
		number: adventofgo.AsInt(name),
		tf:     adventofgo.BinaryToInt(top),
		tb:     adventofgo.BinaryToInt(topBack),
		rf:     adventofgo.BinaryToInt(right),
		rb:     adventofgo.BinaryToInt(rightBack),
		bf:     adventofgo.BinaryToInt(bottom),
		bb:     adventofgo.BinaryToInt(bottomBack),
		lf:     adventofgo.BinaryToInt(left),
		lb:     adventofgo.BinaryToInt(leftBack),
	}
	return t
}

func compareTiles(t, tl tile) int {
	count := 0
	tlSides := []int64{tl.tf, tl.tb, tl.rf, tl.rb, tl.bf, tl.bb, tl.lf, tl.lb}
	if adventofgo.IndexOfInt64(t.tf, tlSides) >= 0 {
		count += 1
	}
	//if adventofgo.IndexOfInt64(t.tb, tlSides) >= 0 {
	//	count += 1
	//}
	if adventofgo.IndexOfInt64(t.rf, tlSides) >= 0 {
		count += 1
	}
	//if adventofgo.IndexOfInt64(t.rb, tlSides) >= 0 {
	//	count += 1
	//}
	if adventofgo.IndexOfInt64(t.bf, tlSides) >= 0 {
		count += 1
	}
	//if adventofgo.IndexOfInt64(t.bb, tlSides) >= 0 {
	//	count += 1
	//}
	if adventofgo.IndexOfInt64(t.lf, tlSides) >= 0 {
		count += 1
	}
	//if adventofgo.IndexOfInt64(t.lb, tlSides) >= 0 {
	//	count += 1
	//}
	return count
}

func isCornerTile(t tile, tiles []tile) bool {
	count := 0
	for _, tl := range tiles {
		if t.number == tl.number {
			continue
		}
		// find how many sides this tile has in common.
		count += compareTiles(t, tl)
	}
	// a corner tile will have 4 of 8 the same.
	return count <= 2

}

func countCorners(input []string) {

	tiles := parseInput(input)
	product := 1
	for _, t := range tiles {
		if isCornerTile(t, tiles) {
			product *= t.number
		}
	}
	fmt.Println("Day 20 part 1")
	fmt.Println(product)

}

func parseInput(input []string) (tiles []tile) {

	pieces := strings.Split(strings.Join(input, "\n"), "\n\n")

	for _, v := range pieces {
		tiles = append(tiles, calculateTile(v))
	}
	return
}
