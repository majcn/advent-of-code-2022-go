package main

import (
	"fmt"
	"regexp"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Node interface {
	Value() float64
}

type ValueNode struct {
	V float64
}

func (node *ValueNode) Value() float64 {
	return node.V
}

type EquationNode struct {
	Left     Node
	Right    Node
	Function func(left float64, right float64) float64
}

func (node *EquationNode) Value() float64 {
	return node.Function(node.Left.Value(), node.Right.Value())
}

type DataType map[string]Node

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	rValueNode := regexp.MustCompile(`^(\w{4}): (\d+)$`)
	rEquationNode := regexp.MustCompile(`^(\w{4}): (\w{4}) ([+\-*/]) (\w{4})$`)

	result := make(map[string]Node, len(data))
	for _, line := range dataSplit {
		match := rValueNode.FindStringSubmatch(line)
		if match != nil {
			result[match[1]] = &ValueNode{V: float64(ParseInt(match[2]))}
		} else {
			result[line[:4]] = &EquationNode{}
		}
	}

	for _, line := range dataSplit {
		match := rEquationNode.FindStringSubmatch(line)
		if match != nil {
			node, _ := result[match[1]].(*EquationNode)
			node.Left = result[match[2]]
			node.Right = result[match[4]]
			switch match[3][0] {
			case '+':
				node.Function = func(left float64, right float64) float64 { return left + right }
			case '-':
				node.Function = func(left float64, right float64) float64 { return left - right }
			case '*':
				node.Function = func(left float64, right float64) float64 { return left * right }
			case '/':
				node.Function = func(left float64, right float64) float64 { return left / right }
			}
		}
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	return int(data["root"].Value())
}

func solvePart2(data DataType) (rc int) {
	rootNode := data["root"].(*EquationNode)
	right := rootNode.Right.Value()

	a := 0
	b := MaxInt / 2
	for {
		c := (a + b) / 2
		data["humn"].(*ValueNode).V = float64(c)
		left := rootNode.Left.Value()

		switch {
		case left > right:
			a = c
		case left < right:
			b = c
		default:
			return c
		}
	}
}

func main() {
	data := parseData(FetchInputData(21))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
