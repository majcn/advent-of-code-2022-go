package filesystem

type fileNode struct {
	size int
}

func (node *fileNode) Size() int {
	return node.size
}

func (node *fileNode) SizeWithoutFolder(withoutFolder Node) int {
	return node.size
}

func (node *fileNode) Walk(visitor func(node Node, isDirectory bool)) {
	visitor(node, false)
}
