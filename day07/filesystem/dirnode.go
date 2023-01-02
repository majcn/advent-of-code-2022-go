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
	for _, childNode := range node.children {
		if childNode != withoutFolder {
			s += childNode.SizeWithoutFolder(withoutFolder)
		}
	}

	return s
}

func (node *dirNode) Walk(visitor func(node Node, isDirectory bool)) {
	for _, childNode := range node.children {
		childNode.Walk(visitor)
	}

	visitor(node, true)
}
