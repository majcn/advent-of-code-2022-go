package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Monkey struct {
	InspectedItems  int
	Items           []int
	Operation       func(x int) int
	TestDivisibleBy int
	TestTrue        int
	TestFalse       int
}

func NewMonkeyFromDescription(description string) Monkey {
	getOperationFunction := func(left string, operator byte, right string) func(int) int {
		if right == "old" {
			return func(x int) int { return x * x }
		}

		rightInt := ParseInt(right)
		switch operator {
		case '+':
			return func(x int) int { return x + rightInt }
		case '*':
			return func(x int) int { return x * rightInt }
		}

		return func(x int) int { return x }
	}

	numbersRegex := regexp.MustCompile(`\d+`)
	functionRegex := regexp.MustCompile(`(old|\d+) ([*+]) (old|\d+)`)

	descriptionLines := strings.Split(description, "\n")

	startingItems := numbersRegex.FindAllString(descriptionLines[1], -1)
	operationParameters := functionRegex.FindStringSubmatch(descriptionLines[2])
	test := numbersRegex.FindString(descriptionLines[3])
	testTrue := numbersRegex.FindString(descriptionLines[4])
	testFalse := numbersRegex.FindString(descriptionLines[5])

	return Monkey{
		InspectedItems:  0,
		Items:           StringsToInts(startingItems),
		Operation:       getOperationFunction(operationParameters[1], operationParameters[2][0], operationParameters[3]),
		TestDivisibleBy: ParseInt(test),
		TestTrue:        ParseInt(testTrue),
		TestFalse:       ParseInt(testFalse),
	}
}

type DataType []Monkey

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n\n")

	result := make([]Monkey, len(dataSplit))
	for i, lines := range dataSplit {
		result[i] = NewMonkeyFromDescription(lines)
	}

	return result
}

func solvePartX(data DataType, rounds int, worryLevelDivisor int) int {
	monkeys := make([]*Monkey, len(data))
	for i, monkey := range data {
		newMonkey := monkey
		monkeys[i] = &newMonkey
	}

	magicNumber := 1
	for _, monkey := range monkeys {
		magicNumber *= monkey.TestDivisibleBy
	}

	for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkey.InspectedItems++
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

	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].InspectedItems > monkeys[j].InspectedItems })
	return monkeys[0].InspectedItems * monkeys[1].InspectedItems
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
