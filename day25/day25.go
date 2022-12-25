package main

import (
	"fmt"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type DataType []string

func parseData(data string) DataType {
	return strings.Split(data, "\n")
}

func encode(n int) string {
	result := []byte{}

	for i := 0; ; i++ {
		result = append(result, '2')
		if decode(string(result)) >= n {
			break
		}
	}

	for i := range result {
		for _, option := range []byte{'=', '-', '0', '1', '2'} {
			result[i] = option
			if decode(string(result)) >= n {
				break
			}
		}
	}

	return string(result)
}

func decode(s string) (rc int) {
	decoder := map[byte]int{'2': 2, '1': 1, '0': 0, '-': -1, '=': -2}

	for i := range s {
		rc += Pow(5, i) * decoder[s[len(s)-i-1]]
	}

	return
}

func solvePart1(data DataType) (rc string) {
	digitalSum := 0
	for _, s := range data {
		digitalSum += decode(s)
	}

	return encode(digitalSum)
}

func solvePart2(data DataType) (rc string) {
	return "Thank you Eric for another wonderful year of AoC!"
}

func main() {
	data := parseData(FetchInputData(23))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
