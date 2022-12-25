package main

import (
	"os"
	"testing"

	. "majcn.si/advent-of-code-2022/util"
)

func ExpectedPart1() int {
	return 15
}

func ExpectedPart2() int {
	return 12
}

var data DataType

func init() {
	input, _ := os.ReadFile("../examples/02.txt")
	data = parseData(string(input))
}

func TestPart1(t *testing.T) {
	AssertTestPartX(t, ExpectedPart1(), solvePart1(data))
}

func TestPart2(t *testing.T) {
	AssertTestPartX(t, ExpectedPart2(), solvePart2(data))
}
