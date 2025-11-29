package utils

import "container/heap"

// ========================
// FIFO QUEUE
// ========================
type Queue[T comparable] []T

func (q *Queue[T]) Push(value T) {
	(*q) = append((*q), value)
}
func (q *Queue[T]) Peek() T {
	return (*q)[0]
}
func (q *Queue[T]) Pop() T {
	popped := (*q)[0]
	(*q) = (*q)[1:]
	return popped
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// ========================
// PRIORITY QUEUE
// ========================
type PriorityQueue[T any] struct {
	items  []T
	sorter func(a, b T) bool
}

func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{sorter: less}
}
func (pq PriorityQueue[T]) Len() int {
	return len(pq.items)
}
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq.sorter(pq.items[i], pq.items[j])
}
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}
func (pq *PriorityQueue[T]) Push(x any) {
	pq.items = append(pq.items, x.(T))
}
func (pq *PriorityQueue[T]) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[:n-1]
	return item
}

// Public Helpers
func (pq *PriorityQueue[T]) Add(item T) {
	heap.Push(pq, item)
}
func (pq *PriorityQueue[T]) Peek() T {
	return pq.items[0]
}
func (pq *PriorityQueue[T]) Remove() T {
	return heap.Pop(pq).(T)
}
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Len() == 0
}
