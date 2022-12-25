package util

type CircualList[T comparable] struct {
	Node *CircualListNode[T]
	Size int
}

type CircualListNode[T comparable] struct {
	Value T
	Prev  *CircualListNode[T]
	Next  *CircualListNode[T]
}

func NewCircualList[T comparable](s []T) CircualList[T] {
	firstNode := &CircualListNode[T]{Value: s[0]}

	node := firstNode
	for _, v := range s[1:] {
		nextNode := &CircualListNode[T]{
			Value: v,
			Prev:  node,
		}

		node.Next = nextNode
		node = nextNode
	}

	node.Next = firstNode
	firstNode.Prev = node

	return CircualList[T]{Node: firstNode, Size: len(s)}
}

func (list *CircualList[T]) AsSlice() []*CircualListNode[T] {
	result := make([]*CircualListNode[T], list.Size)

	node := list.Node
	for i := 0; i < list.Size; i++ {
		result[i] = node
		node = node.Next
	}

	return result
}

func (list *CircualList[T]) FindNode(value T) *CircualListNode[T] {
	node := list.Node
	for i := 0; i < list.Size; i++ {
		if node.Value == value {
			return node
		}

		node = node.Next
	}

	return nil
}

func (list *CircualList[T]) RemoveNode(node *CircualListNode[T]) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
}

func (list *CircualList[T]) AddAfter(after *CircualListNode[T], value T) *CircualListNode[T] {
	node := &CircualListNode[T]{
		Value: value,
		Prev:  after,
		Next:  after.Next,
	}

	after.Next.Prev = node
	after.Next = node

	list.Size++

	return node
}

func (list *CircualList[T]) AddBefore(before *CircualListNode[T], value T) *CircualListNode[T] {
	node := &CircualListNode[T]{
		Value: value,
		Prev:  before.Prev,
		Next:  before,
	}

	before.Prev.Next = node
	before.Prev = node

	list.Size++

	return node
}
