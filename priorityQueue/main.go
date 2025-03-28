package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Name   string
	Expiry int

	Index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest based on expiration number as the priority
	// The lower the expiry, the higher the priority
	return pq[i].Expiry < pq[j].Expiry
}

// We just implement the pre-defined function in interface of heap.

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func main() {

	listItems := []*Item{
		{Name: "Carrot", Expiry: 30},
		{Name: "Potato", Expiry: 45},
		{Name: "Rice", Expiry: 100},
		{Name: "Spinach", Expiry: 5},
	}
	priorityQueue := make(PriorityQueue, len(listItems))

	for i, item := range listItems {
		priorityQueue[i] = item
		priorityQueue[i].Index = i
	}

	heap.Init(&priorityQueue)

	// Print the order by Priority of expiry
	for priorityQueue.Len() > 0 {
		item := heap.Pop(&priorityQueue).(*Item)
		fmt.Printf("Name: %s Expiry: %d\n", item.Name, item.Expiry)
	}
}
