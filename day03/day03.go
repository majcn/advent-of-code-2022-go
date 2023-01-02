package main

import (
	"fmt"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type DataType [][]byte

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make([][]byte, len(dataSplit))
	for i, line := range dataSplit {
		result[i] = []byte(line)
	}

	return result
}

func priority(s byte) int {
	if 'a' <= s && s <= 'z' {
		return int(s) - 96
	} else {
		return int(s) - 38
	}
}

func solvePart1(data DataType) (rc int) {
	for _, line := range data {
		part1Set := NewSet(line[0 : len(line)/2])
		part2Set := NewSet(line[len(line)/2:])
		intersection := part1Set.Intersection(part2Set)
		rc += priority(intersection.Pop())
	}

	return
}

func solvePart2(data DataType) (rc int) {
	for i := 0; i < len(data); i += 3 {
		part1Set := NewSet(data[i])
		part2Set := NewSet(data[i+1])
		part3Set := NewSet(data[i+2])
		intersection := part1Set.Intersection(part2Set)
		intersection = intersection.Intersection(part3Set)
		rc += priority(intersection.Pop())
	}

	return
}

func main() {
	data := parseData(FetchInputData(3))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
