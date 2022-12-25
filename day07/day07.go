package main

import (
	"fmt"
	"strings"

	"majcn.si/advent-of-code-2022/day07/filesystem"
	. "majcn.si/advent-of-code-2022/util"
)

type DataType []string

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "$")

	result := make([]string, len(dataSplit))
	for i, command := range dataSplit {
		result[i] = strings.TrimSpace(command)
	}

	return result
}

func solvePart1(data DataType) (rc int) {
	const MaxSize = 100000

	root := filesystem.BuildFilesystem(data)
	root.Walk(func(node filesystem.Node, isDirectory bool) {
		if !isDirectory {
			return
		}

		size := node.Size()
		if size <= MaxSize {
			rc += size
		}
	})

	return
}

func solvePart2(data DataType) (rc int) {
	const FullSpace, RequiredSpace = 70000000, 30000000
	rc = MaxInt

	root := filesystem.BuildFilesystem(data)
	root.Walk(func(node filesystem.Node, isDirectory bool) {
		if !isDirectory {
			return
		}

		size := root.SizeWithoutFolder(node)
		if size < FullSpace-RequiredSpace {
			rc = Min(rc, node.Size())
		}
	})

	return
}

func main() {
	data := parseData(FetchInputData(7))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
