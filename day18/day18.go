package main

import (
	"fmt"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Point3 struct {
	X int
	Y int
	Z int
}

type DataType []Point3

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make([]Point3, len(dataSplit))
	for i, line := range dataSplit {
		lineSplit := strings.SplitN(line, ",", 3)
		result[i] = Point3{
			X: ParseInt(lineSplit[0]),
			Y: ParseInt(lineSplit[1]),
			Z: ParseInt(lineSplit[2]),
		}
	}

	return result
}

var side1, side2, side3, side4, side5, side6 [4]Point3
var sides [][4]Point3

func init() {
	side1 = [4]Point3{{X: 0, Y: 0, Z: 0}, {X: 1, Y: 0, Z: 0}, {X: 0, Y: 1, Z: 0}, {X: 1, Y: 1, Z: 0}}
	side2 = [4]Point3{{X: 0, Y: 1, Z: 0}, {X: 1, Y: 1, Z: 0}, {X: 0, Y: 1, Z: 1}, {X: 1, Y: 1, Z: 1}}
	side3 = [4]Point3{{X: 0, Y: 0, Z: 1}, {X: 1, Y: 0, Z: 1}, {X: 0, Y: 1, Z: 1}, {X: 1, Y: 1, Z: 1}}
	side4 = [4]Point3{{X: 0, Y: 0, Z: 0}, {X: 1, Y: 0, Z: 0}, {X: 0, Y: 0, Z: 1}, {X: 1, Y: 0, Z: 1}}
	side5 = [4]Point3{{X: 1, Y: 0, Z: 0}, {X: 1, Y: 1, Z: 0}, {X: 1, Y: 0, Z: 1}, {X: 1, Y: 1, Z: 1}}
	side6 = [4]Point3{{X: 0, Y: 0, Z: 0}, {X: 0, Y: 1, Z: 0}, {X: 0, Y: 0, Z: 1}, {X: 0, Y: 1, Z: 1}}
	sides = [][4]Point3{side1, side2, side3, side4, side5, side6}
}

func solvePart1(data DataType) (rc int) {
	cubeSides := make(map[[4]Point3]int)

	for _, side := range sides {
		for _, cube := range data {
			tmp := [4]Point3{}
			for i, point := range side {
				tmp[i] = Point3{
					X: cube.X + point.X,
					Y: cube.Y + point.Y,
					Z: cube.Z + point.Z,
				}
			}

			cubeSides[tmp]++
		}
	}

	for _, c := range cubeSides {
		if c == 1 {
			rc++
		}
	}

	return
}

func findGroup(loc Point3, getNeighborsF func(loc Point3) Set[Point3]) Set[Point3] {
	queue := make(Set[Point3])
	visited := make(Set[Point3])

	queue.Add(loc)
	visited.Add(loc)
	for len(queue) > 0 {
		el := queue.Pop()

		for neighbor := range getNeighborsF(el) {
			if !visited.Contains(neighbor) {
				queue.Add(neighbor)
				visited.Add(neighbor)
			}
		}
	}

	return visited
}

func solvePart2(data DataType) (rc int) {
	cubes := NewSet(data)

	minx, maxx, miny, maxy, minz, maxz := MaxInt, 0, MaxInt, 0, MaxInt, 0
	for cube := range cubes {
		minx = Min(minx, cube.X-1)
		maxx = Max(maxx, cube.X+1)
		miny = Min(miny, cube.Y-1)
		maxy = Max(maxy, cube.Y+1)
		minz = Min(minz, cube.Z-1)
		maxz = Max(maxz, cube.Z+1)
	}

	getNeighborsFOffsets := []Point3{
		{X: 1, Y: 0, Z: 0},
		{X: 0, Y: 1, Z: 0},
		{X: -1, Y: 0, Z: 0},
		{X: 0, Y: -1, Z: 0},
		{X: 0, Y: 0, Z: 1},
		{X: 0, Y: 0, Z: -1},
	}
	getNeighborsF := func(loc Point3) Set[Point3] {
		result := make(Set[Point3], len(getNeighborsFOffsets))

		for _, offset := range getNeighborsFOffsets {
			newLoc := Point3{
				X: loc.X + offset.X,
				Y: loc.Y + offset.Y,
				Z: loc.Z + offset.Z,
			}

			if !cubes.Contains(newLoc) &&
				newLoc.X >= minx && newLoc.X <= maxx &&
				newLoc.Y >= miny && newLoc.Y <= maxy &&
				newLoc.Z >= minz && newLoc.Z <= maxz {

				result.Add(newLoc)
			}
		}

		return result
	}

	outterGroup := findGroup(Point3{X: minx, Y: miny, Z: minz}, getNeighborsF)
	outterGroupCubeSides := make(Set[[4]Point3])
	for _, side := range sides {
		for cube := range outterGroup {
			tmp := [4]Point3{}
			for i, point := range side {
				tmp[i] = Point3{
					X: cube.X + point.X,
					Y: cube.Y + point.Y,
					Z: cube.Z + point.Z,
				}
			}

			outterGroupCubeSides.Add(tmp)
		}
	}

	for _, side := range sides {
		for _, cube := range data {
			tmp := [4]Point3{}
			for i, point := range side {
				tmp[i] = Point3{
					X: cube.X + point.X,
					Y: cube.Y + point.Y,
					Z: cube.Z + point.Z,
				}
			}

			if outterGroupCubeSides.Contains(tmp) {
				rc++
			}
		}
	}

	return
}

func main() {
	data := parseData(FetchInputData(18))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
