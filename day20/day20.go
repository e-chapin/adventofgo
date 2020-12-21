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
	lines []string
}

type coord struct {
	x int
	y int
}

var sideCount = make(map[int64]int)

func main() {

	input := adventofgo.ReadFile("input.txt")
	countCorners(input)
	grid, tiles := makeGrid(input)
	printGrid(grid ,tiles)
}

func rotateTile(t tile) tile {

	// top -> left
	// left -> bottom
	// bottom -> right
	// right -> top

	var newLines []string
	for range t.lines {
		newLines = append(newLines, "")
	}
	for _, l := range t.lines {
		for j, c := range l {
			newLines[len(t.lines)-1-j] += string(c)
		}
	}

	newTile := tile{
		number: t.number,
		tf:     t.rf,
		tb:     t.rb,
		rf:     t.bf,
		rb:     t.bb,
		bf:     t.lf,
		bb:     t.lb,
		lf:     t.tf,
		lb:     t.tb,
		lines:  newLines,
	}
	return newTile
}

func flipTileLeftRight(t tile) tile {

	// top -> top
	// left -> right
	// bottom -> bottom
	// right -> left

	var newLines []string
	for _, line := range t.lines {
		newLines = append(newLines, adventofgo.ReverseString(line))
	}

	newTile := tile{
		number: t.number,
		tf:     t.tb,
		tb:     t.tf,
		rf:     t.lf,
		rb:     t.lb,
		bf:     t.bb,
		bb:     t.bf,
		lf:     t.rf,
		lb:     t.rb,
		lines:  newLines,
	}
	return newTile
}

func flipTileUpDown(t tile) tile {

	// top -> top
	// left -> right
	// bottom -> bottom
	// right -> left

	var newLines []string
	for i := range t.lines {
		newLines = append(newLines, t.lines[len(t.lines)-1-i])
	}

	newTile := tile{
		number: t.number,
		tf:     t.bf,
		tb:     t.bb,
		rf:     t.rb,
		rb:     t.rf,
		bf:     t.tf,
		bb:     t.bb,
		lf:     t.lb,
		lb:     t.lf,
		lines:  newLines,
	}
	return newTile
}

func printGrid(grid map[int][]tile, tiles []tile) {
	var puzzle []string
	for i := 0; i < len(grid)*len(grid[0][0].lines); i += 1 {
		puzzle = append(puzzle, "")
	}

	fixedGrid := make(map[int][]tile)

	// set the first piece in the grid.
	leftCorner := grid[0][0]

	fmt.Println(leftCorner)

	for {
		flippedUpDown := flipTileUpDown(leftCorner)
		if sideCount[leftCorner.tf] == 2 && sideCount[leftCorner.lf] == 2  {
			fixedGrid[0] = append(fixedGrid[0], leftCorner)
			break
		} else if sideCount[flippedUpDown.tf] == 2 && sideCount[flippedUpDown.lf] == 2  {
			fixedGrid[0] = append(fixedGrid[0], flippedUpDown)
			break
		} else {
			leftCorner = rotateTile(leftCorner)
		}
	}

	// fill in the rest of the top row matching the left side to the right side of the previous piece.
	for i:= 1; i < len(grid[0]); i++ {
		piece := grid[0][i]
		for {
			leftPiece := fixedGrid[0][i-1]
			//flippedLeftRight := flipTileLeftRight(piece)
			flippedUpDown := flipTileUpDown(piece)
			if piece.lf == leftPiece.rf {
				fixedGrid[0] = append(fixedGrid[0], piece)
				break
			} else if flippedUpDown.lf == leftPiece.rf {
				fixedGrid[0] = append(fixedGrid[0], flippedUpDown)
				break
			} else {
				piece = rotateTile(piece)
			}
		}
	}

	// fill in the rest of the rows, match up instead of left.

	for row := 1; row < len(grid); row += 1{
		for col := 0; col < len(grid[0]); col += 1 {
			piece := grid[row][col]
			abovePiece := fixedGrid[row-1][col]
			flippedUpDown := flipTileUpDown(piece)
			flippedLeftRight := flipTileLeftRight(piece)
			for {
				if piece.tf == abovePiece.bf {
					fixedGrid[row] = append(fixedGrid[row], piece)
					break
				} else if flippedLeftRight.tf == abovePiece.bf {
					fixedGrid[row] = append(fixedGrid[row], flippedLeftRight)
					break
				} else if flippedUpDown.tf == abovePiece.bf {
					fixedGrid[row] = append(fixedGrid[row], flippedUpDown)
					break
				} else {
					piece = rotateTile(piece)
					flippedUpDown = rotateTile(flippedUpDown)
					flippedLeftRight = rotateTile(flippedLeftRight)
				}
			}
		}


	}

	offset := 0
	for _, piece := range fixedGrid[0] {
		for i, line := range piece.lines {
			puzzle[offset+i] += line
		}
	}
	offset = 8
	for _, piece := range fixedGrid[1] {
		for i, line := range piece.lines {
			puzzle[offset+i] += line
		}
	}
	offset = 16
	for _, piece := range fixedGrid[2] {
		for i, line := range piece.lines {
			puzzle[offset+i] += line
		}
	}
	seamonster := "                  # \n#    ##    ##    ###\n #  #  #  #  #  #   "

	monsterMap := make(map[coord]bool)
	gridMap := make(map[coord]bool)

	// convert seamonster pattern into relative x, y coords
	for y, line := range strings.Split(seamonster, "\n") {
		for x, c := range line {
			if string(c) == "#" {
				monsterMap[coord{x, y}] = true
			}
		}
	}

	fmt.Println(seamonster)

	// 8 possible orientations the grip could be in to find monsters
	correctPuzzleOrientatin := false
	for i := 0; i < 8; i++ {
		for y := range puzzle {
			for x := range puzzle[y] {
				monster := true

				for m := range monsterMap {

					if y+m.y >= len(puzzle) || x+m.x >= len(puzzle[y]) {
						monster = false
						break
					}

					if string(puzzle[y+m.y][x+m.x]) != "1" {
						monster = false
						break
					}
				}
				if monster {
					correctPuzzleOrientatin = true
					for m := range monsterMap {
						gridMap[coord{y+m.y, x+m.x}] = true
					}
				}
			}
		}

		if correctPuzzleOrientatin {
			break
		}

		if i%2 == 0 {
			puzzle = flipPuzzle(puzzle)
			fmt.Println('h')
		} else {
			puzzle = rotatePuzzle(flipPuzzle(puzzle))
			fmt.Println('h')
		}

	}

	// count the 1's not in GridMap
	count := 0
	for y := range puzzle {
		for x := range puzzle[y]{
			if string(puzzle[x][y]) == "1" {
				if gridMap[coord{x, y}]{
					continue
				}
				count += 1
			}
		}
	}
	fmt.Println(count)
}

func flipPuzzle(puzzle []string) []string {
	var newPuzzle []string
	for i := 0; i < len(puzzle); i++ {
		newPuzzle = append(newPuzzle, puzzle[len(puzzle)-1-i])
	}
	return newPuzzle
}

func rotatePuzzle(puzzle []string) []string {
	var newPuzzle []string
	for range puzzle {
		newPuzzle = append(newPuzzle, "")
	}
	for _, l := range puzzle {
		for j, c := range l {
			newPuzzle[len(puzzle)-1-j] += string(c)
		}
	}
	return newPuzzle
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

	// remove borders from the lines for part 2
	lines = lines[1:]
	var newLines []string
	for i, l := range lines {
		if i == 0 || i == len(lines)-1 {
			continue
		}
		newLines = append(newLines, l[1:len(l)-1])
	}

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
		lines:  newLines,
	}

	if _, ok := sideCount[t.tf]; ok {
		sideCount[t.tf] += 1
	} else {
		sideCount[t.tf] = 1
	}

	if _, ok := sideCount[t.tb]; ok {
		sideCount[t.tb] += 1
	} else {
		sideCount[t.tb] = 1
	}
	if _, ok := sideCount[t.rf]; ok {
		sideCount[t.rf] += 1
	} else {
		sideCount[t.rf] = 1
	}
	if _, ok := sideCount[t.rb]; ok {
		sideCount[t.rb] += 1
	} else {
		sideCount[t.rb] = 1
	}
	if _, ok := sideCount[t.bf]; ok {
		sideCount[t.bf] += 1
	} else {
		sideCount[t.bf] = 1
	}
	if _, ok := sideCount[t.bb]; ok {
		sideCount[t.bb] += 1
	} else {
		sideCount[t.bb] = 1
	}
	if _, ok := sideCount[t.lf]; ok {
		sideCount[t.lf] += 1
	} else {
		sideCount[t.lf] = 1
	}
	if _, ok := sideCount[t.lb]; ok {
		sideCount[t.lb] += 1
	} else {
		sideCount[t.lb] = 1
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

func RemoveIndex(t []tile, index int) []tile {
	return append(t[:index], t[index+1:]...)
}

// this is gross, but loop through each row and column to find all the corners and edge pieces,
// then fill in the rest.
func makeGrid(input []string) (map[int][]tile, []tile){
	tiles := parseInput(input)
	var corners []tile
	var grid = make(map[int][]tile)
	row := 0
	// find corners
	var cornerNumbers []int
	for _, t := range tiles {
		if isCornerTile(t, tiles) {
			corners = append(corners, t)
			cornerNumbers = append(cornerNumbers, t.number)
		}
	}

	//Pick a random one to be the top left?
	grid[row] = append(grid[row], corners[0])
	corners = corners[1:]
	// Add tiles to the right of this one, until another corners side fits and completes the row.
	// After this for loop grid[0] will be the top of the puzzle.
	for {
		foundCorner := false
		for i, c := range corners {
			if compareTiles(c, grid[row][len(grid[row])-1]) == 1 {
				grid[row] = append(grid[row], c)
				corners = RemoveIndex(corners, i)
				foundCorner = true
			}
		}
		if foundCorner {
			break
		}

		//otherwise check pieces for a fit. I think each edge only has one matching edge.
		for i, t := range tiles {
			if t.number == grid[row][len(grid[row])-1].number {
				continue
			}
			if compareTiles(t, grid[row][len(grid[row])-1]) == 1{
				grid[row] = append(grid[row], t)
				tiles = RemoveIndex(tiles, i)
				break
			}
		}
	}

	// now build the left side of the puzzle, starting from the top left corner again.
	// Each piece needs to make a new index in the grid.
	var index = 1
	for {
		// start at 1 since grid at index 0 is done

		foundCorner := false
		for i, c := range corners {
			if compareTiles(c, grid[index-1][0]) == 1 {
				grid[index] = append(grid[index], c)
				corners = RemoveIndex(corners, i)
				index += 1
				foundCorner = true
			}
		}
		if foundCorner {
			break
		}

		//otherwise check pieces for a fit. I think each edge only has one matching edge.
		for i, t := range tiles {
			if t.number == grid[index-1][0].number {
				continue
			}
			if compareTiles(t, grid[index-1][0]) == 1{
				grid[index] = append(grid[index], t)
				index += 1
				tiles = RemoveIndex(tiles, i)
				break
			}
		}
	}

	// build the bottom row. Index should be the bottom row.
	for {
		row = index-1
		foundCorner := false
		for i, c := range corners {
			if compareTiles(c, grid[row][len(grid[row])-1]) == 1 {
				grid[row] = append(grid[row], c)
				corners = RemoveIndex(corners, i)
				foundCorner = true
			}
		}
		if foundCorner {
			break
		}

		//otherwise check pieces for a fit. I think each edge only has one matching edge.
		for i, t := range tiles {
			if t.number == grid[row][len(grid[row])-1].number {
				continue
			}
			if compareTiles(t, grid[row][len(grid[row])-1]) == 1{
				grid[row] = append(grid[row], t)
				tiles = RemoveIndex(tiles, i)
				break
			}
		}
	}
	// now fill in any middle rows.

	// build the bottom row. Index should be the bottom row.
	for row = 1; row < len(grid)-1; row++ {
		index = 0
		if len(grid[row]) == 1 {
			index = 1
		}
		for {
			// grid[0] is guaranteed a full row, so stop once this row is the same size.
			if index >= len(grid[0]) {
				break
			}

			// no more corners so find pieces that match this piece.
			// my assumption is that each side is unique, so if the piece fits the piece to its left
			// it should also fit the piece above and below.
			for i, t := range tiles {
				if t.number == grid[row][len(grid[row])-1].number {
					continue
				}
				if adventofgo.IndexOfInt(t.number, cornerNumbers) >= 0 {
					continue
				}
				if compareTiles(t, grid[row][len(grid[row])-1]) == 1 {
					grid[row] = append(grid[row], t)
					tiles = RemoveIndex(tiles, i)
					index += 1
					break
				}
			}
		}
	}
	// should now have one row of tiles, but not necessarily oriented correctly
	return grid, tiles
}

func countCorners(input []string) {

	tiles := parseInput(input)
	product := 1
	for _, t := range tiles {
		if isCornerTile(t, tiles) {
			fmt.Println("Corner tile", t.number)
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
