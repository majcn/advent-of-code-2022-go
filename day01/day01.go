package main

import (
	"fmt"
	"sort"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type DataType [][]int

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n\n")

	result := make([][]int, len(dataSplit))
	for i, lines := range dataSplit {
		result[i] = StringsToInts(strings.Split(lines, "\n"))
	}

	return result
}

func solvePartX(data DataType) []int {
	result := make([]int, len(data))
	for i, group := range data {
		result[i] = Sum(group)
	}
	return result
}

func solvePart1(data DataType) (rc int) {
	return Max(solvePartX(data)...)
}

func solvePart2(data DataType) (rc int) {
	sumOfGroups := solvePartX(data)
	sort.Sort(sort.Reverse(sort.IntSlice(sumOfGroups)))
	return Sum(sumOfGroups[:3])
}

func main() {
	data := parseData(FetchInputData(1))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
