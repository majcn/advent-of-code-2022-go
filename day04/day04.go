package main

import (
	"fmt"
	"regexp"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Range struct {
	Min int
	Max int
}

type State struct {
	P1 Range
	P2 Range
}

type DataType []State

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	r := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)

	result := make([]State, len(dataSplit))
	for i, line := range dataSplit {
		match := r.FindStringSubmatch(line)
		p1 := Range{Min: ParseInt(match[1]), Max: ParseInt(match[2])}
		p2 := Range{Min: ParseInt(match[3]), Max: ParseInt(match[4])}
		result[i] = State{P1: p1, P2: p2}
	}

	return result
}

type PartResult struct {
	S1           Set[int]
	S2           Set[int]
	Intersection Set[int]
}

func solvePartX(data DataType) []PartResult {
	result := make([]PartResult, len(data))
	for i, state := range data {
		s1 := make(Set[int], state.P1.Max-state.P1.Min+1)
		for i := state.P1.Min; i <= state.P1.Max; i++ {
			s1.Add(i)
		}

		s2 := make(Set[int], state.P2.Max-state.P2.Min+1)
		for i := state.P2.Min; i <= state.P2.Max; i++ {
			s2.Add(i)
		}

		intersection := s1.Intersection(&s2)

		result[i] = PartResult{
			S1:           s1,
			S2:           s2,
			Intersection: intersection,
		}
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	for _, s := range solvePartX(data) {
		if len(s.Intersection) == Min(len(s.S1), len(s.S2)) {
			rc++
		}
	}

	return
}

func solvePart2(data DataType) (rc int) {
	for _, s := range solvePartX(data) {
		if len(s.Intersection) > 0 {
			rc++
		}
	}

	return
}

func main() {
	data := parseData(FetchInputData(4))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
