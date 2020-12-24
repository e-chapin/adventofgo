package main

import (
	"adventofgo"
	"fmt"
)

// inspiration for coord system: https://www.redblobgames.com/grids/hexagons/
type hexCoord struct {
	x int
	y int
	z int
}

var coords = make(map[hexCoord]bool)

func placeTile(input []string) {
	//var coords = make(map[hexCoord]bool)
	for _, v := range input {
		x, y, z := 0, 0, 0
		for i := 0; i < len(v); i += 1 {
			c := string(v[i])

			if c == "s" || c == "n" {
				i += 1
				c += string(v[i])
			}
			fmt.Println("checking c", c)
			switch c {
			case "e":
				x += 1
				y -= 1
			case "se":
				z += 1
				y -= 1
			case "sw":
				z += 1
				x -= 1
			case "w": // west
				x -= 1
				y += 1
			case "nw":
				y += 1
				z -= 1
			case "ne":
				x += 1
				z -= 1
			}
			coord := hexCoord{x,y, z}
			fmt.Println(coord)
			if _, ok := coords[coord]; !ok {
				coords[coord] = false // set initial to white, it will get flipped on next line
			}
			if i == len(v)-1{
				coords[coord] = !coords[coord] // true is black up
			}

		}
	}
	count := 0
	for _, v := range coords {
		if v {
			count += 1
		}
	}
	fmt.Println(count)
	return
}

func cloneCoords() map[hexCoord]bool {

	clone := make(map[hexCoord]bool)
	for c, v := range coords {
		clone[c] = v
	}

	return clone
}

func checkNeighbors(c hexCoord) int {

	count := 0
	var neighbors []hexCoord
	neighbors = append(neighbors, hexCoord{c.x+1, c.y-1, c.z})
	neighbors = append(neighbors, hexCoord{c.x, c.y-1, c.z+1})
	neighbors = append(neighbors, hexCoord{c.x-1, c.y, c.z+1})
	neighbors = append(neighbors, hexCoord{c.x-1, c.y+1, c.z})
	neighbors = append(neighbors, hexCoord{c.x, c.y+1, c.z-1})
	neighbors = append(neighbors, hexCoord{c.x+1, c.y, c.z-1})

	for _, n := range neighbors {
		if _, ok := coords[n]; ok {
			if coords[n] {
				count += 1
			}
		}
	}

	return count
}


func doArt(days int){

	for i := 1; i <= days; i += 1 {
		var clone = cloneCoords()
		for c, v := range coords {
			// true is black
			count := checkNeighbors(c)
			if v {
				if count == 0 || count > 2 {
					clone[c] = false
				}
			} else {
				if count == 2 {
					clone[c] = true
				}
			}

			// some catch all to make sure outside hex's are flipped
			// when two outermost adjecent in my map turn black
			// too tired think of a smart way to do this.
			for x := -2; x < 2; x += 1 {
				for y := -2; y < 2; y += 1 {
					for z := -2; z < 2; z += 1 {
						outer := hexCoord{c.x+x, c.y+y, c.z+z}
						if _, ok := coords[outer]; !ok {
							count = checkNeighbors(outer)
							if count == 2 {
								clone[outer] = true
							}
						}
					}
				}
			}
		}
		coords = clone
	}
	count := 0
	for _, v := range coords {
		if v {
			count += 1
		}
	}
	fmt.Println(count)
}

func main() {

	input := adventofgo.ReadFile("input.txt")
	fmt.Println("Day 24 Part 1")
	placeTile(input)
	fmt.Println("Day 24 Part 2")
	doArt(100)

}
