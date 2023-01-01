package main

import (
	"fmt"
	"regexp"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Command struct {
	Move            int
	ChangeDirection byte
}

type Location Point
type Direction Point

type DataType struct {
	StartLocation  Location
	StartDirection Direction
	Grid           Set[Location]
	GridSize       Location
	Wall           Set[Location]
	Commands       []Command
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	grid := make(Set[Location])
	wall := make(Set[Location])
	startLocation := Location{X: -1, Y: -1}
	for y, line := range dataSplit[:len(dataSplit)-2] {
		for x, v := range line {
			switch v {
			case '.':
				grid.Add(Location{X: x, Y: y})
				if startLocation.X < 0 {
					startLocation = Location{X: x, Y: y}
				}
			case '#':
				wall.Add(Location{X: x, Y: y})
			}
		}
	}

	commandLine := dataSplit[len(dataSplit)-1]
	commandMatches := regexp.MustCompile(`(\d+)([LR])`).FindAllStringSubmatch(commandLine, -1)
	commands := make([]Command, len(commandMatches)*2+1)
	for i, match := range commandMatches {
		commands[i*2] = Command{Move: ParseInt(match[1])}
		commands[i*2+1] = Command{ChangeDirection: match[2][0]}
	}
	commands[len(commands)-1] = Command{
		Move: ParseInt(regexp.MustCompile(`\d+$`).FindString(commandLine)),
	}

	return DataType{
		StartLocation:  startLocation,
		StartDirection: directionRight,
		Grid:           grid,
		GridSize:       Location{X: len(dataSplit[0]), Y: len(dataSplit) - 2},
		Wall:           wall,
		Commands:       commands,
	}
}

type NextDirectionMapKey struct {
	FromDirection Direction
	Instruction   byte
}

var directionLeft, directionRight, directionUp, directionDown Direction
var nextDirectionMap map[NextDirectionMapKey]Direction

func init() {
	directionLeft = Direction{X: -1, Y: 0}
	directionRight = Direction{X: 1, Y: 0}
	directionUp = Direction{X: 0, Y: -1}
	directionDown = Direction{X: 0, Y: 1}

	nextDirectionMap = map[NextDirectionMapKey]Direction{
		{FromDirection: directionLeft, Instruction: 'L'}: directionDown,
		{FromDirection: directionLeft, Instruction: 'R'}: directionUp,

		{FromDirection: directionRight, Instruction: 'L'}: directionUp,
		{FromDirection: directionRight, Instruction: 'R'}: directionDown,

		{FromDirection: directionUp, Instruction: 'L'}: directionLeft,
		{FromDirection: directionUp, Instruction: 'R'}: directionRight,

		{FromDirection: directionDown, Instruction: 'L'}: directionRight,
		{FromDirection: directionDown, Instruction: 'R'}: directionLeft,
	}
}

func calculateScore(location Location, direction Direction) int {
	var facingScore int
	switch direction {
	case directionRight:
		facingScore = 0
	case directionDown:
		facingScore = 1
	case directionLeft:
		facingScore = 2
	case directionUp:
		facingScore = 3
	}

	return 4*(location.X+1) + 1000*(location.Y+1) + facingScore
}

func solvePartX(data DataType, nextStateF func(location Location, direction Direction) (Location, Direction)) int {
	startLocation, startDirection, grid, commands := data.StartLocation, data.StartDirection, data.Grid, data.Commands

	myLocation := startLocation
	myDirection := startDirection
	for _, c := range commands {
		if c.ChangeDirection == 'L' || c.ChangeDirection == 'R' {
			myDirection = nextDirectionMap[NextDirectionMapKey{FromDirection: myDirection, Instruction: c.ChangeDirection}]
			continue
		}

		for i := 0; i < c.Move; i++ {
			newLocation, newDirection := nextStateF(myLocation, myDirection)
			if !grid.Contains(newLocation) {
				break
			}

			myLocation, myDirection = newLocation, newDirection
		}
	}

	return calculateScore(myLocation, myDirection)
}

func nextStatePart1(grid Set[Location], gridSize Location, wall Set[Location], location Location, direction Direction) (Location, Direction) {
	result := Location{X: location.X + direction.X, Y: location.Y + direction.Y}

	for !wall.Contains(result) && !grid.Contains(result) {
		result.X += direction.X
		result.Y += direction.Y

		switch {
		case result.X >= gridSize.X:
			result.X = 0
		case result.X < 0:
			result.X = gridSize.X - 1
		case result.Y >= gridSize.Y:
			result.Y = 0
		case result.Y < 0:
			result.Y = gridSize.Y - 1
		}
	}

	return result, direction
}

func solvePart1(data DataType) (rc int) {
	grid, gridSize, wall := data.Grid, data.GridSize, data.Wall

	return solvePartX(data, func(location Location, direction Direction) (Location, Direction) {
		return nextStatePart1(grid, gridSize, wall, location, direction)
	})
}

func nextStatePart2(location Location, direction Direction) (Location, Direction) {
	switch {
	case location.X == 50 && 0 <= location.Y && location.Y <= 49 && direction == directionLeft:
		return Location{X: 0, Y: -location.Y + 149}, directionRight
	case location.Y == 0 && 50 <= location.X && location.X <= 99 && direction == directionUp:
		return Location{X: 0, Y: location.X + 100}, directionRight
	case location.Y == 0 && 100 <= location.X && location.X <= 149 && direction == directionUp:
		return Location{X: location.X - 100, Y: 199}, directionUp
	case location.X == 149 && 0 <= location.Y && location.Y <= 49 && direction == directionRight:
		return Location{X: 99, Y: 149 - location.Y}, directionLeft
	case location.Y == 49 && 100 <= location.X && location.X <= 149 && direction == directionDown:
		return Location{X: 99, Y: location.X - 50}, directionLeft
	case location.X == 99 && 50 <= location.Y && location.Y <= 99 && direction == directionRight:
		return Location{X: location.Y + 50, Y: 49}, directionUp
	case location.X == 50 && 50 <= location.Y && location.Y <= 99 && direction == directionLeft:
		return Location{X: location.Y - 50, Y: 100}, directionDown
	case location.X == 99 && 100 <= location.Y && location.Y <= 149 && direction == directionRight:
		return Location{X: 149, Y: -location.Y + 149}, directionLeft
	case location.Y == 149 && 50 <= location.X && location.X <= 99 && direction == directionDown:
		return Location{X: 49, Y: location.X + 100}, directionLeft
	case location.X == 49 && 150 <= location.Y && location.Y <= 199 && direction == directionRight:
		return Location{X: location.Y - 100, Y: 149}, directionUp
	case location.Y == 199 && 0 <= location.X && location.X <= 49 && direction == directionDown:
		return Location{X: location.X + 100, Y: 0}, directionDown
	case location.X == 0 && 150 <= location.Y && location.Y <= 199 && direction == directionLeft:
		return Location{X: location.Y - 100, Y: 0}, directionDown
	case location.X == 0 && 100 <= location.Y && location.Y <= 149 && direction == directionLeft:
		return Location{X: 50, Y: -location.Y + 149}, directionRight
	case location.Y == 100 && 0 <= location.X && location.X <= 49 && direction == directionUp:
		return Location{X: 50, Y: location.X + 50}, directionRight
	default:
		return Location{X: location.X + direction.X, Y: location.Y + direction.Y}, direction
	}
}

func solvePart2(data DataType) (rc int) {
	return solvePartX(data, func(location Location, direction Direction) (Location, Direction) {
		return nextStatePart2(location, direction)
	})
}

func main() {
	data := parseData(FetchInputData(22))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
