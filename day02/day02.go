package main

import (
	"fmt"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type State struct {
	P1 byte
	P2 byte
}

type DataType []State

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make([]State, len(dataSplit))
	for i, line := range dataSplit {
		p1, p2, _ := strings.Cut(line, " ")

		result[i] = State{P1: p1[0], P2: p2[0]}
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	// A = Rock, B = Paper, C = Scissors, X = Rock, Y = Paper, Z = Scissors
	score := map[State]int{
		{P1: 'A', P2: 'X'}: 3 + 1,
		{P1: 'A', P2: 'Y'}: 6 + 2,
		{P1: 'A', P2: 'Z'}: 0 + 3,
		{P1: 'B', P2: 'X'}: 0 + 1,
		{P1: 'B', P2: 'Y'}: 3 + 2,
		{P1: 'B', P2: 'Z'}: 6 + 3,
		{P1: 'C', P2: 'X'}: 6 + 1,
		{P1: 'C', P2: 'Y'}: 0 + 2,
		{P1: 'C', P2: 'Z'}: 3 + 3,
	}

	for _, state := range data {
		rc += score[state]
	}

	return
}

func solvePart2(data DataType) (rc int) {
	// A = Rock, B = Paper, C = Scissors, X = Lose, Y = Draw, Z = Win
	score := map[State]int{
		{P1: 'A', P2: 'X'}: 0 + 3,
		{P1: 'A', P2: 'Y'}: 3 + 1,
		{P1: 'A', P2: 'Z'}: 6 + 2,
		{P1: 'B', P2: 'X'}: 0 + 1,
		{P1: 'B', P2: 'Y'}: 3 + 2,
		{P1: 'B', P2: 'Z'}: 6 + 3,
		{P1: 'C', P2: 'X'}: 0 + 2,
		{P1: 'C', P2: 'Y'}: 3 + 3,
		{P1: 'C', P2: 'Z'}: 6 + 1,
	}

	for _, state := range data {
		rc += score[state]
	}

	return
}

func main() {
	data := parseData(FetchInputData(2))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
