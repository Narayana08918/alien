package pq

import (
	"container/heap"

	"github.com/Narayana08918/alien/models"
)

type items []*models.City

type PriorityQueue struct {
	queue *items
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{queue: &items{}}

	heap.Init(pq.queue)
	return pq
}

func (pq *PriorityQueue) Push(s *models.City) {
	heap.Push(pq.queue, s)
}

func (pq *PriorityQueue) Pop() *models.City {
	return heap.Pop(pq.queue).(*models.City)
}

func (pq *PriorityQueue) Size() int {
	return pq.queue.Len()
}

func (pq items) Len() int {
	return len(pq)
}

func (pq items) Less(i, j int) bool {
	return len(pq[i].Out) > len(pq[j].Out)
}

func (pq items) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *items) Push(x interface{}) {
	item := x.(*models.City)
	*pq = append(*pq, item)
}

func (pq *items) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
