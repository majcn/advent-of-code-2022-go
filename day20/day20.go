package main

import (
	"fmt"
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type DataType []int

func parseData(data string) DataType {
	return StringsToInts(strings.Split(data, "\n"))
}

func solvePartX(data DataType, nrMixes int) (rc int) {
	myList := NewCircualList(data)
	myListAsSlice := myList.AsSlice()

	for n := 0; n < nrMixes; n++ {
		for i := range myListAsSlice {
			node := myListAsSlice[i]
			nodeValue := node.Value

			switch {
			case nodeValue > 0:
				myNextNode := node.Next

				myList.RemoveNode(node)
				node = nil

				repeat := (nodeValue % (len(myListAsSlice) - 1)) - 1
				for c := 0; c < repeat; c++ {
					myNextNode = myNextNode.Next
				}

				myListAsSlice[i] = myList.AddAfter(myNextNode, nodeValue)
			case nodeValue < 0:
				myPrevNode := node.Prev

				myList.RemoveNode(node)
				node = nil

				repeat := ((-1 * nodeValue) % (len(myListAsSlice) - 1)) - 1
				for c := 0; c < repeat; c++ {
					myPrevNode = myPrevNode.Prev
				}

				myListAsSlice[i] = myList.AddBefore(myPrevNode, nodeValue)
			}
		}
	}

	node := myList.FindNode(0)
	for i := 1; i <= 3000; i++ {
		node = node.Next
		if i%1000 == 0 {
			rc += node.Value
		}
	}

	return
}

func solvePart1(data DataType) (rc int) {
	return solvePartX(data, 1)
}

func solvePart2(data DataType) (rc int) {
	modifiedData := make(DataType, len(data))
	for i, v := range data {
		modifiedData[i] = v * 811589153
	}

	return solvePartX(modifiedData, 10)
}

func main() {
	data := parseData(FetchInputData(20))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
