package main

import (
	"fmt"

	. "majcn.si/advent-of-code-2022/util"
)

type DataType []byte

func parseData(data string) DataType {
	return []byte(data)
}

func solvePartX(data DataType, n int) int {
	for i := 0; ; i++ {
		if NewSet(data[i:i+n]).Len() == n {
			return i + n
		}
	}
}

func solvePart1(data DataType) (rc int) {
	return solvePartX(data, 4)
}

func solvePart2(data DataType) (rc int) {
	return solvePartX(data, 14)
}

func main() {
	data := parseData(FetchInputData(6))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
