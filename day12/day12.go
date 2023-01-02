package main

import (
	"fmt"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Node struct {
	X         int
	Y         int
	Mark      byte
	Neighbors []*Node
}

type DataType struct {
	StartNode *Node
	EndNode   *Node
	Nodes     map[Point]*Node
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	numberOfNodes := len(dataSplit) * len(dataSplit[0])
	result := DataType{Nodes: make(map[Point]*Node, numberOfNodes)}
	for y, line := range dataSplit {
		for x, v := range line {
			var node *Node
			switch v {
			case 'S':
				node = &Node{X: x, Y: y, Mark: 'a'}
				result.StartNode = node
			case 'E':
				node = &Node{X: x, Y: y, Mark: 'z'}
				result.EndNode = node
			default:
				node = &Node{X: x, Y: y, Mark: byte(v)}
			}
			result.Nodes[Point{X: x, Y: y}] = node
		}
	}

	for location, node := range result.Nodes {
		node.Neighbors = make([]*Node, 0, 4)
		for _, neighbour := range GetNeighbors4() {
			newLocation := location.Add(neighbour)
			if nNode, ok := result.Nodes[newLocation]; ok && int(nNode.Mark)-int(node.Mark) < 2 {
				node.Neighbors = append(node.Neighbors, nNode)
			}
		}
	}

	return result
}

func aStarSearchAlgorithm(startNode *Node, endNode *Node) int {
	h := func(n *Node) int { return Pow(endNode.X-n.X, 2) + Pow(endNode.Y-n.Y, 2) }
	d := func(c *Node, n *Node) int { return 1 }

	openSet := NewSet([]*Node{startNode})
	queue := NewPriorityQueue(startNode)

	gScore := map[*Node]int{startNode: 0}
	for queue.Len() > 0 {
		current, score := queue.Pop()

		if current == endNode {
			return score
		}

		openSet.Remove(current)
		for _, neighbour := range current.Neighbors {
			tentativeGScore := gScore[current] + d(current, neighbour)
			if gScoreValue, ok := gScore[neighbour]; ok && tentativeGScore >= gScoreValue {
				continue
			}

			gScore[neighbour] = tentativeGScore
			fScore := tentativeGScore + h(neighbour)
			if openSet.Contains(neighbour) {
				queue.Fix(neighbour, fScore)
			} else {
				openSet.Add(neighbour)
				queue.Push(neighbour, fScore)
			}
		}
	}

	return MaxInt
}

func solvePart1(data DataType) (rc int) {
	startNode, endNode := data.StartNode, data.EndNode
	return aStarSearchAlgorithm(startNode, endNode)
}

func solvePart2(data DataType) (rc int) {
	endNode, nodes := data.EndNode, data.Nodes
	rc = MaxInt

	for _, node := range nodes {
		if node.Mark == 'a' {
			score := aStarSearchAlgorithm(node, endNode)
			rc = Min(rc, score)
		}
	}

	return
}

func main() {
	data := parseData(FetchInputData(12))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
