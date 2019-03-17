package main

import (
	"container/heap"
	"fmt"
)

func main() {
	items := IntHeaps{{1}, {10}, {11}, {13}, {456}, {970}, {500}}
	//sort.Sort(items)
	heap.Init(&items)
	fmt.Println(items)
}

type IntHeap struct {
	t     int
	//index int
}

type IntHeaps []IntHeap

func (h IntHeaps) Len() int {
	return len(h)
}
func (h IntHeaps) Less(i, j int) bool {
	return h[i].t < h[j].t
}
func (h IntHeaps) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	//h[i].index = i
	//h[j].index = j
}

func (h *IntHeaps) Push(x interface{}) {
	//n := len(*h)
	item := x.(*IntHeap)
	//item.index = n
	*h = append(*h, *item)
}

func (h *IntHeaps) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	//item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}
