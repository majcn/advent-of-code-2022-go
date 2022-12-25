package filesystem

type dirNode struct {
	parent   Node
	children map[string]Node
}

func (node *dirNode) Size() int {
	return node.SizeWithoutFolder(nil)
}

func (node *dirNode) SizeWithoutFolder(withoutFolder Node) int {
	s := 0
	for childName := range node.children {
		childNode := node.children[childName]
		if childNode != withoutFolder {
			s += childNode.SizeWithoutFolder(withoutFolder)
		}
	}

	return s
}

func (node *dirNode) Walk(visitor func(node Node, isDirectory bool)) {
	for child := range node.children {
		node.children[child].Walk(visitor)
	}

	visitor(node, true)
}
