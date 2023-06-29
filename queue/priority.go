package queue

import (
	"container/heap"

	"github.com/dhm116/maze-generator/queue/internal"
)

type PriorityQueue struct {
	queue   internal.PriorityQueue
	itemMap map[any]*internal.Item
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{
		queue:   make(internal.PriorityQueue, 0),
		itemMap: make(map[any]*internal.Item),
	}
	heap.Init(&pq.queue)
	return pq
}

func (pq *PriorityQueue) Len() int {
	return pq.queue.Len()
}

func (pq *PriorityQueue) Push(value any, priority int) {
	i := &internal.Item{
		Value:    value,
		Priority: priority,
	}
	pq.itemMap[value] = i
	heap.Push(&pq.queue, i)
}

func (pq *PriorityQueue) Pop() any {
	i := heap.Pop(&pq.queue).(*internal.Item)
	delete(pq.itemMap, i.Value)
	return i.Value
}

func (pq *PriorityQueue) Update(value any, priority int) {
	i, ok := pq.itemMap[value]
	if !ok {
		pq.Push(value, priority)
		return
	}
	pq.queue.Update(i, value, priority)
}
