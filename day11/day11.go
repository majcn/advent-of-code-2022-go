package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Monkey struct {
	Items           []int
	Operation       func(x int) int
	TestDivisibleBy int
	TestTrue        int
	TestFalse       int
}

func (monkey *Monkey) SetOperation(left string, operator byte, right string) {
	if right == "old" {
		monkey.Operation = func(x int) int { return x * x }
		return
	}

	rightInt := ParseInt(right)
	switch operator {
	case '+':
		monkey.Operation = func(x int) int { return x + rightInt }
	case '*':
		monkey.Operation = func(x int) int { return x * rightInt }
	}
}

type DataType []*Monkey

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	numbersRegex := regexp.MustCompile(`\d+`)
	functionRegex := regexp.MustCompile(`(old|\d+) ([*+]) (old|\d+)`)

	result := make([]*Monkey, (len(dataSplit)+1)/7)
	for i := 0; i < len(dataSplit); i += 7 {
		startingItems := numbersRegex.FindAllString(dataSplit[i+1], -1)
		operationParameters := functionRegex.FindStringSubmatch(dataSplit[i+2])
		test := numbersRegex.FindString(dataSplit[i+3])
		testTrue := numbersRegex.FindString(dataSplit[i+4])
		testFalse := numbersRegex.FindString(dataSplit[i+5])

		monkey := &Monkey{}
		monkey.Items = StringsToInts(startingItems)
		monkey.SetOperation(operationParameters[1], operationParameters[2][0], operationParameters[3])
		monkey.TestDivisibleBy = ParseInt(test)
		monkey.TestTrue = ParseInt(testTrue)
		monkey.TestFalse = ParseInt(testFalse)

		result[i/7] = monkey
	}

	return result
}

func copyMonkeys(monkeys []*Monkey) []*Monkey {
	result := make([]*Monkey, len(monkeys))
	for i, monkey := range monkeys {
		newMonkey := *monkey
		result[i] = &newMonkey
	}
	return result
}

func solvePartX(data DataType, rounds int, worryLevelDivisor int) int {
	monkeys := copyMonkeys(data)

	magicNumber := 1
	for _, monkey := range monkeys {
		magicNumber *= monkey.TestDivisibleBy
	}

	result := make([]int, len(monkeys))
	for round := 0; round < rounds; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.Items {
				result[i]++
				newItem := (monkey.Operation(item) / worryLevelDivisor) % magicNumber
				if newItem%monkey.TestDivisibleBy == 0 {
					monkeys[monkey.TestTrue].Items = append(monkeys[monkey.TestTrue].Items, newItem)
				} else {
					monkeys[monkey.TestFalse].Items = append(monkeys[monkey.TestFalse].Items, newItem)
				}
			}
			monkey.Items = []int{}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result[0] * result[1]
}

func solvePart1(data DataType) (rc int) {
	return solvePartX(data, 20, 3)
}

func solvePart2(data DataType) (rc int) {
	return solvePartX(data, 10000, 1)
}

func main() {
	data := parseData(FetchInputData(11))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
