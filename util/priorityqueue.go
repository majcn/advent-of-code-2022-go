package util

import "container/heap"

type PriorityQueue[T comparable] struct {
	queue *priorityQueueInternal
}

func NewPriorityQueue[T comparable](initItem T) PriorityQueue[T] {
	pq := PriorityQueue[T]{queue: &priorityQueueInternal{}}
	pq.Push(initItem, 0)
	return pq
}

func (pq *PriorityQueue[T]) Len() int {
	return len(*pq.queue)
}

func (pq *PriorityQueue[T]) Pop() (T, int) {
	item := heap.Pop(pq.queue).(*priorityQueueItemInternal)
	return item.value.(T), item.score
}

func (pq *PriorityQueue[T]) Push(value T, score int) {
	item := &priorityQueueItemInternal{value: value, score: score}
	heap.Push(pq.queue, item)
}

func (pq *PriorityQueue[T]) Fix(value T, newScore int) {
	for i, node := range *pq.queue {
		if node.value == value {
			node.score = newScore
			heap.Fix(pq.queue, i)
			return
		}
	}
}

type priorityQueueItemInternal struct {
	value any
	score int
	index int
}

type priorityQueueInternal []*priorityQueueItemInternal

func (pq priorityQueueInternal) Len() int {
	return len(pq)
}

func (pq priorityQueueInternal) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq priorityQueueInternal) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueueInternal) Push(x any) {
	n := len(*pq)
	item := x.(*priorityQueueItemInternal)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueueInternal) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
