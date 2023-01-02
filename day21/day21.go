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
	v float64
}

func (node ValueNode) Value() float64 {
	return node.v
}

func NewValueNode(value int) ValueNode {
	return ValueNode{v: float64(value)}
}

type EquationNode struct {
	nodeMapper map[string]Node
	left       string
	right      string
	function   func(left float64, right float64) float64
}

func (node EquationNode) LeftNode() Node {
	return node.nodeMapper[node.left]
}

func (node EquationNode) RightNode() Node {
	return node.nodeMapper[node.right]
}

func (node EquationNode) Value() float64 {
	return node.function(node.nodeMapper[node.left].Value(), node.nodeMapper[node.right].Value())
}

func NewEquationNodeFromDescription(description string, nodeMapper map[string]Node) EquationNode {
	r := regexp.MustCompile(`^\w{4}: (\w{4}) ([+\-*/]) (\w{4})$`)

	getFunction := func(operand byte) func(left float64, right float64) float64 {
		switch operand {
		case '+':
			return func(left float64, right float64) float64 { return left + right }
		case '-':
			return func(left float64, right float64) float64 { return left - right }
		case '*':
			return func(left float64, right float64) float64 { return left * right }
		case '/':
			return func(left float64, right float64) float64 { return left / right }
		}

		return func(left, right float64) float64 { return 0 }
	}

	match := r.FindStringSubmatch(description)

	return EquationNode{
		nodeMapper: nodeMapper,
		left:       match[1],
		right:      match[3],
		function:   getFunction(match[2][0]),
	}
}

type DataType map[string]Node

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	rValueNode := regexp.MustCompile(`^\w{4}: (\d+)$`)

	result := make(map[string]Node, len(data))
	for _, line := range dataSplit {
		match := rValueNode.FindStringSubmatch(line)
		if match != nil {
			result[line[:4]] = NewValueNode(ParseInt(match[1]))
		} else {
			result[line[:4]] = NewEquationNodeFromDescription(line, result)
		}
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	return int(data["root"].Value())
}

func solvePart2(data DataType) (rc int) {
	rootNode := data["root"].(EquationNode)
	right := rootNode.RightNode().Value()

	a := 0
	b := MaxInt / 2
	for {
		c := (a + b) / 2
		data["humn"] = NewValueNode(c)
		left := rootNode.LeftNode().Value()

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
