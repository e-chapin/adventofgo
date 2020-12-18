package main

import (
	"adventofgo"
	"fmt"
	"time"
)


type coord struct {
	x int
	y int
	z int
	w int
}

var cubeSize int
var max = 30 // some generic large enough number?

func main() {

	strlines := adventofgo.ReadFile("input.txt")
	cubeSize = len(strlines)
	coords := initCoords(cubeSize)
	// setup initial coords
	for i, row := range strlines {
		z := max / 2
		for j, val := range  row {
			// add z (half of max size) to get this 2d slice right in the middle of space
			c := coord{i+z, j+z, z, z}
			active := val == '#'
			coords[c] = active
		}
	}

	for i:= 0; i < 6; i++ {
		coords = doCycle(coords)
	}
	start := time.Now()
	fmt.Println("Day 17 Part 1")
	print(countActive(coords))
	fmt.Println(time.Since(start))

}

func initCoords(size int) map[coord]bool {
	coords := make(map[coord]bool)
	for z := 0; z < max; z++ {
		for x:= 0; x < max; x++ {
			for y := 0; y < max; y++ {
				for w:= 0; w < max; w++ {
					c := coord{x, y, z, w}
					coords[c] = false
				}

			}
		}
	}
	return coords
}

func cloneSpace(coords map[coord]bool) map[coord]bool {
	tmpCoords := make(map[coord]bool)
	for c, active := range coords {
		clone := coord{c.x, c.y, c.z, c.w}
		tmpCoords[clone] = active
	}
	return tmpCoords
}

func checkNearbyCubes(coords map[coord]bool, c  coord) int{

	var cx, cy, cz, cw int

	activeNeighbors := 0
	inspectedCoords := 0
	// gotta be a more clever way to do this
	// edit, part 2: Now there for sure has to be some maths I don't know to
	// help calculate these. I think the key is that it starts small so its predictable?
	for _, x := range []int{-1, 0, 1} {
		cx = c.x +x
		for _,  y := range []int{-1, 0, 1} {
			cy = c.y + y
			for _, z := range []int{-1, 0, 1} {
				cz = c.z + z
				for _, w := range []int{-1, 0, 1} {
					cw = c.w + w
					if cx == c.x && cy == c.y && cz == c.z && cw == c.w{
						continue
					}
					neighborCoord := coord{cx, cy, cz, cw}
					if  coords[neighborCoord] {
						activeNeighbors += 1
					}
					inspectedCoords += 1
				}
			}
		}
	}
	return activeNeighbors
}

func doCycle(coords map[coord]bool) map[coord]bool {

	newCoords := cloneSpace(coords)
	for c, a := range coords {
		activeNeighbors := checkNearbyCubes(coords, c)
		// this coord is active, stay active if 2 or 3 neighbors are
		if a {
			if activeNeighbors != 2 && activeNeighbors != 3 {
				newCoords[c] = false
			}
			// if this coord is inactive, become active if exactly 3 neighbors are active.
		} else {
			if activeNeighbors == 3 {
				newCoords[c] = true
			}
		}
	}
	return newCoords
}

func countActive(coords map[coord]bool) int {
	active := 0
	for _, v := range coords {
		if v {
			active += 1
		}
	}
	return active
}