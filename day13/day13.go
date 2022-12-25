package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type DataType [][]any

func parseData(data string) DataType {
	dataSplit := strings.Split(strings.ReplaceAll(data, "\n\n", "\n"), "\n")

	result := make([][]any, len(dataSplit))
	for i, v := range dataSplit {
		json.Unmarshal([]byte(v), &result[i])
	}

	return DataType(result)
}

func compare(left []any, right []any) int {
	size := Min(len(left), len(right))

	for i := 0; i < size; i++ {
		leftList, isLeftList := left[i].([]any)
		rightList, isRightList := right[i].([]any)
		leftNum, isLeftNum := left[i].(float64)
		rightNum, isRightNum := right[i].(float64)

		var r int
		switch {
		case isLeftList && isRightList:
			r = compare(leftList, rightList)
		case isLeftNum && isRightNum:
			r = int(rightNum - leftNum)
		case isLeftList:
			r = compare(leftList, []any{rightNum})
		default:
			r = compare([]any{leftNum}, rightList)
		}

		if r != 0 {
			return r
		}
	}

	return len(right) - len(left)
}

func solvePart1(data DataType) (rc int) {
	for i := 0; i < len(data); i += 2 {
		if compare(data[i], data[i+1]) > 0 {
			rc += (i/2 + 1)
		}
	}
	return
}

func solvePart2(data DataType) (rc int) {
	divider1 := "[[2]]"
	divider2 := "[[6]]"
	result := make([]int, 0, 2)

	var divider1Unmarshal, divider2Unmarshal []any
	json.Unmarshal([]byte(divider1), &divider1Unmarshal)
	json.Unmarshal([]byte(divider2), &divider2Unmarshal)

	packages := append(data, divider1Unmarshal, divider2Unmarshal)
	sort.Slice(packages, func(i, j int) bool { return compare(packages[i], packages[j]) > 0 })
	for i, v := range packages {
		sv := fmt.Sprint(v)
		if sv == divider1 || sv == divider2 {
			result = append(result, i+1)
			if len(result) == 2 {
				break
			}
		}
	}

	return result[0] * result[1]
}

func main() {
	data := parseData(FetchInputData(13))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
