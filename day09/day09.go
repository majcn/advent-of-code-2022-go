package main

import (
	"fmt"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Command struct {
	Direction byte
	Steps     int
}

type DataType []Command

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make([]Command, len(dataSplit))
	for i, line := range dataSplit {
		direction, steps, _ := strings.Cut(line, " ")
		result[i] = Command{
			Direction: direction[0],
			Steps:     ParseInt(steps),
		}
	}

	return result
}

func solvePartX(data DataType, ropeSize int) int {
	rope := make([]Point, ropeSize)

	visited := NewSet([]Point{{X: 0, Y: 0}})
	for _, command := range data {
		for s := 0; s < command.Steps; s++ {
			switch command.Direction {
			case 'R':
				rope[0] = rope[0].Add(Point{X: 1, Y: 0})
			case 'D':
				rope[0] = rope[0].Add(Point{X: 0, Y: 1})
			case 'L':
				rope[0] = rope[0].Add(Point{X: -1, Y: 0})
			case 'U':
				rope[0] = rope[0].Add(Point{X: 0, Y: -1})
			}

			for i := 1; i < ropeSize; i++ {
				head := rope[i-1]
				tail := &rope[i]

				if Abs(head.X-tail.X) > 1 || Abs(head.Y-tail.Y) > 1 {
					var diffX int
					switch {
					case head.X < tail.X:
						diffX = -1
					case head.X > tail.X:
						diffX = 1
					default:
						diffX = 0
					}

					var diffY int
					switch {
					case head.Y < tail.Y:
						diffY = -1
					case head.Y > tail.Y:
						diffY = 1
					default:
						diffY = 0
					}

					*tail = tail.Add(Point{X: diffX, Y: diffY})
				}
			}

			visited.Add(rope[len(rope)-1])
		}
	}

	return len(visited)
}

func solvePart1(data DataType) (rc int) {
	return solvePartX(data, 2)
}

func solvePart2(data DataType) (rc int) {
	return solvePartX(data, 10)
}

func main() {
	data := parseData(FetchInputData(9))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
