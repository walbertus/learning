package leetcode

import (
    "container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(i, j int) bool {
    return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    value := old[len(old)-1]
    *h = old[:len(old)-1]
    return value
}

// https://leetcode.com/problems/last-stone-weight/
func lastStoneWeight(stones []int) int {
    intHeap := IntHeap(stones)
    heap.Init(&intHeap)
    for intHeap.Len() > 1 {
        first := heap.Pop(&intHeap).(int)
        second := heap.Pop(&intHeap).(int)
        heap.Push(&intHeap, first - second)
    }
    if intHeap.Len() == 0 {
        return 0
    }
    return heap.Pop(&intHeap).(int)
}
