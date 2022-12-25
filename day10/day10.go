package main

import (
	"fmt"
	"strings"

	"majcn.si/advent-of-code-2022/day10/interpreter"
	. "majcn.si/advent-of-code-2022/util"
)

type DataType []string

func parseData(data string) DataType {
	return DataType(strings.Split(data, "\n"))
}

func solvePart1(data DataType) (rc int) {
	p := interpreter.New(data)

	for !p.End {
		switch p.Cycle + 1 {
		case 20, 60, 100, 140, 180, 220:
			rc += (p.Cycle + 1) * p.RegisterX
		}

		p.ExecSingleCycle()
	}

	return
}

func solvePart2(data DataType) (rc string) {
	display := [6][40]byte{}
	for y, line := range display {
		for x := range line {
			display[y][x] = ' '
		}
	}

	p := interpreter.New(data)

	for !p.End {
		x := p.Cycle % 40
		y := p.Cycle / 40
		if x >= p.RegisterX-1 && x <= p.RegisterX+1 {
			display[y][x] = '#'
		}

		p.ExecSingleCycle()
	}

	var sb strings.Builder
	sb.Grow(len(display) * (len(display[0]) + 1))
	for _, line := range display {
		sb.Write(line[:])
		sb.WriteByte('\n')
	}

	return sb.String()
}

func main() {
	data := parseData(FetchInputData(10))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
