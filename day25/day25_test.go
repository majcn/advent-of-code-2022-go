package main

import (
	"testing"

	. "majcn.si/advent-of-code-2022/util"
)

func ExpectedPart1() string {
	return "2-=2-0=-0-=0200=--21"
}

func ExpectedPart2() string {
	return "Thank you Eric for another wonderful year of AoC!"
}

var data DataType

func init() {
	data = InitData(parseData)
}

func TestPart1(t *testing.T) {
	AssertTestPartX(t, ExpectedPart1(), solvePart1(data))
}

func TestPart2(t *testing.T) {
	AssertTestPartX(t, ExpectedPart2(), solvePart2(data))
}
