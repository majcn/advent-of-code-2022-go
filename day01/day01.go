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
	for j, lines := range dataSplit {
		lineSplit := strings.Split(lines, "\n")

		result[j] = make([]int, len(lineSplit))
		for i, v := range lineSplit {
			result[j][i] = ParseInt(v)
		}
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
