package leetcode

import "container/heap"

const MAX_SIZE = 1000002

// TODO: Optimize by only compute things that are queried rather than precompute everything
func maxPoints(grid [][]int, queries []int) []int {
	score := make([]int, MAX_SIZE)

	currentValue := 0

	isVisited := newSet()
	currentScore := 0

	pq := newPQ()
	heap.Init(&pq)
	heap.Push(&pq, []int{grid[0][0], 0, 0})
	isVisited.Add(0, 0)

	for pq.Len() > 0 {
		pqValue := heap.Pop(&pq).([]int)
		gridValue, xGrid, yGrid := pqValue[0], pqValue[1], pqValue[2]
		if gridValue >= currentValue {
			for i := currentValue; i <= gridValue; i++ {
				score[i] = currentScore
			}
			currentValue = gridValue + 1
		}
		currentScore++

		if xGrid > 0 && !isVisited.Has(xGrid-1, yGrid) {
			heap.Push(&pq, []int{grid[xGrid-1][yGrid], xGrid - 1, yGrid})
			isVisited.Add(xGrid-1, yGrid)
		}
		if xGrid < len(grid)-1 && !isVisited.Has(xGrid+1, yGrid) {
			heap.Push(&pq, []int{grid[xGrid+1][yGrid], xGrid + 1, yGrid})
			isVisited.Add(xGrid+1, yGrid)
		}
		if yGrid > 0 && !isVisited.Has(xGrid, yGrid-1) {
			heap.Push(&pq, []int{grid[xGrid][yGrid-1], xGrid, yGrid - 1})
			isVisited.Add(xGrid, yGrid-1)
		}
		if yGrid < len(grid[0])-1 && !isVisited.Has(xGrid, yGrid+1) {
			heap.Push(&pq, []int{grid[xGrid][yGrid+1], xGrid, yGrid + 1})
			isVisited.Add(xGrid, yGrid+1)
		}
	}

	for i := currentValue; i < MAX_SIZE; i++ {
		score[i] = currentScore
	}

	answer := make([]int, 0)
	for _, query := range queries {
		answer = append(answer, score[query])
	}

	return answer
}

type PQ struct {
	// This is lower priority queue
	// first value is cell value, second and third are coordinates
	container [][]int
}

func newPQ() PQ {
	return PQ{
		container: make([][]int, 0),
	}
}

func (pq PQ) Len() int { return len(pq.container) }

func (pq PQ) Less(i, j int) bool {
	return pq.container[i][0] < pq.container[j][0]
}

func (pq *PQ) Pop() interface{} {
	old := pq.container
	n := len(old)
	item := old[n-1]
	pq.container = old[0 : n-1]
	return item
}

func (pq *PQ) Push(x interface{}) {
	pq.container = append(pq.container, x.([]int))
}

func (pq PQ) Swap(i, j int) {
	pq.container[i], pq.container[j] = pq.container[j], pq.container[i]
}

type Set struct {
	container map[int]map[int]bool
}

func newSet() *Set {
	return &Set{
		container: make(map[int]map[int]bool),
	}
}

func (s *Set) Add(x, y int) {
	if _, ok := s.container[x]; !ok {
		s.container[x] = make(map[int]bool)
	}
	s.container[x][y] = true
}

func (s *Set) Has(x, y int) bool {
	if _, ok := s.container[x]; !ok {
		return false
	}
	return s.container[x][y]
}
