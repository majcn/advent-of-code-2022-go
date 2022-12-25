package main

import (
	"fmt"
	"regexp"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type DataType [][]Location

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	r := regexp.MustCompile(`(\d+),(\d+)`)

	result := make([][]Location, len(dataSplit))
	for i, line := range dataSplit {
		matches := r.FindAllStringSubmatch(line, -1)
		result[i] = make([]Location, len(matches))
		for j, match := range matches {
			result[i][j] = Location{X: ParseInt(match[1]), Y: ParseInt(match[2])}
		}
	}

	return result
}

func buildGrid(data [][]Location) Set[Location] {
	grid := make(Set[Location])
	for _, path := range data {
		for i := 0; i < len(path)-1; i++ {
			start := path[i]
			end := path[i+1]
			startX, endX := Min(start.X, end.X), Max(start.X, end.X)
			startY, endY := Min(start.Y, end.Y), Max(start.Y, end.Y)
			for x := startX; x <= endX; x++ {
				for y := startY; y <= endY; y++ {
					grid.Add(Location{X: x, Y: y})
				}
			}
		}
	}
	return grid
}

func solvePartX(data DataType, endPredicate func(sand Location, maxY int) bool) int {
	grid := buildGrid(data)
	maxY := 0
	for location := range grid {
		maxY = Max(maxY, location.Y)
	}

	locationDown := Location{X: 0, Y: 1}
	locationDownLeft := Location{X: -1, Y: 1}
	locationDownRight := Location{X: 1, Y: 1}

	for i := 0; ; i++ {
		sand := Location{X: 500, Y: 0}
		for sand.Y != maxY+1 {
			nextSand := sand.Add(locationDown)
			if grid.Contains(nextSand) {
				nextSand = sand.Add(locationDownLeft)
			}
			if grid.Contains(nextSand) {
				nextSand = sand.Add(locationDownRight)
			}
			if grid.Contains(nextSand) {
				break
			}
			sand = nextSand
		}

		if endPredicate(sand, maxY) {
			return i
		}

		grid.Add(sand)
	}
}

func solvePart1(data DataType) (rc int) {
	return solvePartX(data, func(sand Location, maxY int) bool { return sand.Y == maxY+1 })
}

func solvePart2(data DataType) (rc int) {
	goal := Location{X: 500, Y: 0}
	return solvePartX(data, func(sand Location, _ int) bool { return sand == goal }) + 1
}

func main() {
	data := parseData(FetchInputData(14))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
