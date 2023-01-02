package main

import (
	"fmt"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type DataType [][]int

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make([][]int, len(dataSplit))
	for y, line := range dataSplit {
		result[y] = make([]int, len(line))
		for x, v := range line {
			result[y][x] = ParseInt(v)
		}
	}

	return result
}

type Range struct {
	From int
	To   int
	Step int
}

func solvePart1(data DataType) (rc int) {
	lenX := len(data[0])
	lenY := len(data)

	visible := make(Set[Point])

	for y := 0; y < lenY; y++ {
		maxTree := -1
		for x := 0; x < lenX; x++ {
			if data[y][x] > maxTree {
				visible.Add(Point{X: x, Y: y})
				maxTree = data[y][x]
			}
		}
	}

	for y := 0; y < lenY; y++ {
		maxTree := -1
		for x := lenX - 1; x >= 0; x-- {
			if data[y][x] > maxTree {
				visible.Add(Point{X: x, Y: y})
				maxTree = data[y][x]
			}
		}
	}

	for x := 0; x < lenX; x++ {
		maxTree := -1
		for y := 0; y < lenY; y++ {
			if data[y][x] > maxTree {
				visible.Add(Point{X: x, Y: y})
				maxTree = data[y][x]
			}
		}
	}

	for x := 0; x < lenX; x++ {
		maxTree := -1
		for y := lenY - 1; y >= 0; y-- {
			if data[y][x] > maxTree {
				visible.Add(Point{X: x, Y: y})
				maxTree = data[y][x]
			}
		}
	}

	return visible.Len()
}

func solvePart2(data DataType) (rc int) {
	lenX := len(data[0])
	lenY := len(data)

	for y, line := range data {
		for x, myTree := range line {
			result := 1

			c := 0
			for xx := x + 1; xx < lenX; xx++ {
				c++
				if myTree <= data[y][xx] {
					break
				}
			}
			result *= c

			c = 0
			for xx := x - 1; xx >= 0; xx-- {
				c++
				if myTree <= data[y][xx] {
					break
				}
			}
			result *= c

			c = 0
			for yy := y + 1; yy < lenY; yy++ {
				c++
				if myTree <= data[yy][x] {
					break
				}
			}
			result *= c

			c = 0
			for yy := y - 1; yy >= 0; yy-- {
				c++
				if myTree <= data[yy][x] {
					break
				}
			}
			result *= c

			rc = Max(rc, result)
		}
	}

	return
}

func main() {
	data := parseData(FetchInputData(8))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
