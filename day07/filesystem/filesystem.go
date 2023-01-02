package filesystem

import (
	"strings"

	. "majcn.si/advent-of-code-2022/util"
)

type Node interface {
	Size() int
	SizeWithoutFolder(withoutFolder Node) int
	Walk(visitor func(node Node, isDirectory bool))
}

func cmdCommand(command string, root Node, currentNode Node) Node {
	_, dirName, _ := strings.Cut(command, " ")
	switch dirName {
	case "/":
		return root
	case "..":
		return currentNode.(*dirNode).parent
	default:
		return currentNode.(*dirNode).children[dirName]
	}
}

func lsCommand(command string, root Node, currentNode Node) {
	lines := strings.Split(command, "\n")[1:]
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "dir"):
			_, dirName, _ := strings.Cut(line, " ")
			childNode := &dirNode{
				parent:   currentNode,
				children: map[string]Node{},
			}
			currentNode.(*dirNode).children[dirName] = childNode
		default:
			fileSize, fileName, _ := strings.Cut(line, " ")
			fileNode := fileNode{
				size: ParseInt(fileSize),
			}
			currentNode.(*dirNode).children[fileName] = &fileNode
		}
	}
}

func BuildFilesystem(commands []string) Node {
	root := Node(&dirNode{
		children: make(map[string]Node),
	})

	currentNode := root
	for _, command := range commands {
		switch {
		case strings.HasPrefix(command, "cd"):
			currentNode = cmdCommand(command, root, currentNode)
		case strings.HasPrefix(command, "ls"):
			lsCommand(command, root, currentNode)
		}
	}

	return root
}
