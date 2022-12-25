package main

import (
	"fmt"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type DataType Set[Location]

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make(Set[Location])
	for y, line := range dataSplit {
		for x, v := range line {
			if v == '#' {
				result.Add(Location{X: x, Y: y})
			}
		}
	}

	return DataType(result)
}

func getNextLocation(elfs Set[Location], elf Location, priorities []byte) Location {
	hasNoNeighbours := true
	for _, neighbour := range GetNeighbours8() {
		newLocation := elf.Add(neighbour)
		if elfs.Contains(newLocation) {
			hasNoNeighbours = false
			break
		}
	}

	if hasNoNeighbours {
		return elf
	}

	for _, priority := range priorities {
		switch priority {
		case 'N':
			if !elfs.Contains(elf.Add(Location{X: -1, Y: -1})) && !elfs.Contains(elf.Add(Location{X: 0, Y: -1})) && !elfs.Contains(elf.Add(Location{X: 1, Y: -1})) {
				return elf.Add(Location{X: 0, Y: -1})
			}
		case 'S':
			if !elfs.Contains(elf.Add(Location{X: -1, Y: 1})) && !elfs.Contains(elf.Add(Location{X: 0, Y: 1})) && !elfs.Contains(elf.Add(Location{X: 1, Y: 1})) {
				return elf.Add(Location{X: 0, Y: 1})
			}
		case 'W':
			if !elfs.Contains(elf.Add(Location{X: -1, Y: -1})) && !elfs.Contains(elf.Add(Location{X: -1, Y: 0})) && !elfs.Contains(elf.Add(Location{X: -1, Y: 1})) {
				return elf.Add(Location{X: -1, Y: 0})
			}
		case 'E':
			if !elfs.Contains(elf.Add(Location{X: 1, Y: -1})) && !elfs.Contains(elf.Add(Location{X: 1, Y: 0})) && !elfs.Contains(elf.Add(Location{X: 1, Y: 1})) {
				return elf.Add(Location{X: 1, Y: 0})
			}
		}
	}

	return elf
}

func solvePartX(data DataType, goal func(int, bool) bool) (int, Set[Location]) {
	priorities := [][]byte{
		{'N', 'S', 'W', 'E'},
		{'S', 'W', 'E', 'N'},
		{'W', 'E', 'N', 'S'},
		{'E', 'N', 'S', 'W'},
	}

	elfs := make(Set[Location], len(data))
	for el := range data {
		elfs.Add(el)
	}

	for i := 0; ; i++ {
		blockedLocationsSet := make(Set[Location])
		nextLocations := make(map[Location]Location)
		for elf := range elfs {
			nextElfLocation := getNextLocation(elfs, elf, priorities[i%4])
			if elf == nextElfLocation {
				continue
			}

			if blockedLocationsSet.Contains(nextElfLocation) {
				continue
			}

			if _, ok := nextLocations[nextElfLocation]; ok {
				blockedLocationsSet.Add(nextElfLocation)
				delete(nextLocations, nextElfLocation)
				continue
			}

			nextLocations[nextElfLocation] = elf
		}

		for nextLocation, prevLocation := range nextLocations {
			elfs.Remove(prevLocation)
			elfs.Add(nextLocation)
		}

		if goal(i+1, len(nextLocations) == 0) {
			return i + 1, elfs
		}
	}
}

func solvePart1(data DataType) (rc int) {
	_, elfs := solvePartX(data, func(i int, _ bool) bool { return i == 10 })

	minx, maxx, miny, maxy := MaxInt, 0, MaxInt, 0
	for elf := range elfs {
		minx = Min(minx, elf.X)
		maxx = Max(maxx, elf.X)
		miny = Min(miny, elf.Y)
		maxy = Max(maxy, elf.Y)
	}

	return (maxy-miny+1)*(maxx-minx+1) - len(elfs)
}

func solvePart2(data DataType) (rc int) {
	rc, _ = solvePartX(data, func(_ int, nothingHappens bool) bool { return nothingHappens })

	return
}

func main() {
	data := parseData(FetchInputData(23))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
