package leetcode

import (
	"container/heap"
)

func putMarbles(weights []int, k int) int64 {
	if k == 1 || len(weights) == k {
		return 0
	}

	maxCutPoints := newPQ()
	minCutPoints := newPQ()
	heap.Init(maxCutPoints)
	heap.Init(minCutPoints)

	// cutPoints := make([]int, 0)
	for i := 1; i < len(weights); i++ {
		// cutPoints = append(cutPoints, weights[i]+weights[i-1])
		heap.Push(maxCutPoints, int64(weights[i]+weights[i-1]))
		if maxCutPoints.Len() > k-1 {
			heap.Pop(maxCutPoints)
		}
		heap.Push(minCutPoints, int64(-(weights[i] + weights[i-1])))
		if minCutPoints.Len() > k-1 {
			heap.Pop(minCutPoints)
		}
	}

	// sort.Ints(cutPoints)

	minScore := int64(weights[0] + weights[len(weights)-1])
	maxScore := int64(weights[0] + weights[len(weights)-1])

	for i := 0; i < k-1; i++ {
		minScore -= heap.Pop(minCutPoints).(int64)
		maxScore += heap.Pop(maxCutPoints).(int64)
		// minScore += int64(cutPoints[i])
		// maxScore += int64(cutPoints[len(cutPoints)-1-i])
	}

	return maxScore - minScore
}

type PQ struct {
	container []int64
}

func newPQ() *PQ {
	return &PQ{
		container: make([]int64, 0),
	}
}

func (pq *PQ) Len() int {
	return len(pq.container)
}

func (pq *PQ) Less(i, j int) bool {
	return pq.container[i] < pq.container[j]
}

func (pq *PQ) Swap(i, j int) {
	pq.container[i], pq.container[j] = pq.container[j], pq.container[i]
}

func (pq *PQ) Push(x interface{}) {
	pq.container = append(pq.container, x.(int64))
}

func (pq *PQ) Pop() interface{} {
	x := pq.container[len(pq.container)-1]
	pq.container = pq.container[:len(pq.container)-1]
	return x
}
