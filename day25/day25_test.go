package main

import (
	"os"
	"testing"

	. "majcn.si/advent-of-code-2022/util"
)

func ExpectedPart1() string {
	return "2=-1=0"
}

func ExpectedPart2() string {
	return "⭐️⭐️"
}

var data DataType

func init() {
	input, _ := os.ReadFile("../examples/25.txt")
	data = parseData(string(input))
}

func TestPart1(t *testing.T) {
	AssertTestPartX(t, ExpectedPart1(), solvePart1(data))
}

func TestPart2(t *testing.T) {
	AssertTestPartX(t, ExpectedPart2(), solvePart2(data))
}
