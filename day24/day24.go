package main

import (
	"fmt"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Blizzard struct {
	Point
	Direction byte
}

type DataType struct {
	Blizzards     []Blizzard
	BlizzardsBBox BBox
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := DataType{
		Blizzards: []Blizzard{},
		BlizzardsBBox: BBox{
			MinX: 1,
			MaxX: len(dataSplit[0]) - 2,
			MinY: 1,
			MaxY: len(dataSplit) - 2,
		},
	}

	for y, line := range dataSplit {
		for x, v := range line {
			switch v {
			case '^', 'v', '<', '>':
				result.Blizzards = append(result.Blizzards, Blizzard{Point: Point{X: x, Y: y}, Direction: byte(v)})
			}
		}
	}

	return result
}

type GraphNode struct {
	Point
	BlizzardsIndex int
}

type QueueNode struct {
	GraphNode
	Time int
}

func bfs(startNode GraphNode, endLocation Point, getNeighborsF func(node GraphNode, startLocation, endLocation Point) []GraphNode) (int, GraphNode) {
	queue := make([]QueueNode, 0)
	visited := make(Set[GraphNode])

	queue = append(queue, QueueNode{GraphNode: startNode, Time: 0})
	visited.Add(startNode)
	for len(queue) > 0 {
		queueEl := queue[0]
		queue = queue[1:]

		node := queueEl.GraphNode

		if node.Point == endLocation {
			return queueEl.Time, queueEl.GraphNode
		}

		for _, neighbor := range getNeighborsF(node, startNode.Point, endLocation) {
			if !visited.Contains(neighbor) {
				queue = append(queue, QueueNode{GraphNode: neighbor, Time: queueEl.Time + 1})
				visited.Add(neighbor)
			}
		}
	}

	return MaxInt, GraphNode{}
}

func getNeighbors(node GraphNode, blizzardsBBox BBox, getBlizzardLocations func(i int) Set[Point], startLocation Point, endLocation Point) []GraphNode {
	result := make([]GraphNode, 0, 5)

	neighbors := append(GetNeighbors4(), Point{X: 0, Y: 0})

	newBlizzardsIndex := node.BlizzardsIndex + 1
	newBlizzards := getBlizzardLocations(newBlizzardsIndex)
	for _, neighbor := range neighbors {
		newLocation := node.Point.Add(neighbor)

		if newLocation == endLocation {
			return []GraphNode{{Point: newLocation, BlizzardsIndex: newBlizzardsIndex}}
		}

		if !newBlizzards.Contains(newLocation) {
			if newLocation == startLocation || newLocation.InBBox(blizzardsBBox) {
				result = append(result, GraphNode{
					Point:          newLocation,
					BlizzardsIndex: newBlizzardsIndex,
				})
			}
		}
	}

	return result
}

func getBlizzardLocations(blizzardsBBox BBox, blizzards []Blizzard, time int) Set[Point] {
	xBlizzardPath := blizzardsBBox.MaxX - blizzardsBBox.MinX + 1
	yBlizzardPath := blizzardsBBox.MaxY - blizzardsBBox.MinY + 1

	result := make(Set[Point])
	for _, blizzard := range blizzards {
		switch blizzard.Direction {
		case '^':
			result.Add(Point{X: blizzard.X, Y: Mod(blizzard.Y-1-time, yBlizzardPath) + 1})
		case 'v':
			result.Add(Point{X: blizzard.X, Y: Mod(blizzard.Y-1+time, yBlizzardPath) + 1})
		case '<':
			result.Add(Point{X: Mod(blizzard.X-1-time, xBlizzardPath) + 1, Y: blizzard.Y})
		case '>':
			result.Add(Point{X: Mod(blizzard.X-1+time, xBlizzardPath) + 1, Y: blizzard.Y})
		}
	}

	return result
}

func solvePartX(data DataType, paths []Point) (rc int) {
	blizzards, blizzardsBBox := data.Blizzards, data.BlizzardsBBox
	allBlizzardLocations := make(map[int]Set[Point])

	getBlizzardLocationsF := func(i int) Set[Point] {
		if _, ok := allBlizzardLocations[i]; ok {
			return allBlizzardLocations[i]
		}

		blizzardLocations := getBlizzardLocations(blizzardsBBox, blizzards, i)
		allBlizzardLocations[i] = blizzardLocations
		return blizzardLocations
	}

	getNeighborsF := func(node GraphNode, startLocation, endLocation Point) []GraphNode {
		return getNeighbors(node, blizzardsBBox, getBlizzardLocationsF, startLocation, endLocation)
	}

	currentBlizzardIndex := 0
	node := GraphNode{Point: paths[0], BlizzardsIndex: currentBlizzardIndex}
	for _, p := range paths[1:] {
		time, goalNode := bfs(node, p, getNeighborsF)

		node = goalNode
		rc += time
	}

	return
}

func solvePart1(data DataType) (rc int) {
	blizzardsBBox := data.BlizzardsBBox

	path := []Point{
		{X: 1, Y: 0},
		{X: blizzardsBBox.MaxX, Y: blizzardsBBox.MaxY + 1},
	}

	return solvePartX(data, path)
}

func solvePart2(data DataType) (rc int) {
	blizzardsBBox := data.BlizzardsBBox

	path := []Point{
		{X: 1, Y: 0},
		{X: blizzardsBBox.MaxX, Y: blizzardsBBox.MaxY + 1},
		{X: 1, Y: 0},
		{X: blizzardsBBox.MaxX, Y: blizzardsBBox.MaxY + 1},
	}

	return solvePartX(data, path)
}

func main() {
	data := parseData(FetchInputData(24))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
