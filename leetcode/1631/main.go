package leetcode

import (
    "container/heap"
    "math"
)

type Point struct {
    Row, Column int
}

func minimumEffortPath(heights [][]int) int {
    row := len(heights)
    column := len(heights[0])
    dist := map[int]map[int]int{}
    for i := 0; i < len(heights); i++ {
        dist[i] = map[int]int{}
    }
    dist[0][0] = 0
    pq := &PointHeap{}
    heap.Init(pq)
    heap.Push(pq, PointDistance{
        Point: Point{0, 0},
        Cost:  0,
    })
    for len(*pq) > 0 {
        current := heap.Pop(pq).(PointDistance)
        currentPoint := current.Point
        currentCost := current.Cost
        if currentCost != dist[currentPoint.Row][currentPoint.Column] {
            continue
        }
        topPoint := Point{
            Row:    currentPoint.Row - 1,
            Column: currentPoint.Column,
        }
        if isValidPoint(topPoint, row, column) {
            cost := moveCost(currentPoint, topPoint, heights)
            addWhenSmaller(topPoint, max(cost, currentCost), dist, pq)
        }
        leftPoint := Point{
            Row:    currentPoint.Row,
            Column: currentPoint.Column - 1,
        }
        if isValidPoint(leftPoint, row, column) {
            cost := moveCost(currentPoint, leftPoint, heights)
            addWhenSmaller(leftPoint, max(cost, currentCost), dist, pq)
        }
        rightPoint := Point{
            Row:    currentPoint.Row,
            Column: currentPoint.Column + 1,
        }
        if isValidPoint(rightPoint, row, column) {
            cost := moveCost(currentPoint, rightPoint, heights)
            addWhenSmaller(rightPoint, max(cost, currentCost), dist, pq)
        }
        bottomPoint := Point{
            Row:    currentPoint.Row + 1,
            Column: currentPoint.Column,
        }
        if isValidPoint(bottomPoint, row, column) {
            cost := moveCost(currentPoint, bottomPoint, heights)
            addWhenSmaller(bottomPoint, max(cost, currentCost), dist, pq)
        }
    }
    return dist[row-1][column-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func isValidPoint(point Point, row, column int) bool {
    if point.Row < 0 || point.Column < 0 || point.Row >= row || point.Column >= column {
        return false
    }
    return true
}

func moveCost(a, b Point, heights [][]int) int {
    return abs(heights[a.Row][a.Column] - heights[b.Row][b.Column])
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func addWhenSmaller(point Point, cost int, dist map[int]map[int]int, pq *PointHeap) {
    currentCost := math.MaxInt32
    if val, ok := dist[point.Row][point.Column]; ok {
        currentCost = val
    }
    if cost >= currentCost {
        return
    }
    dist[point.Row][point.Column] = cost
    heap.Push(pq, PointDistance{
        Point: point,
        Cost:  cost,
    })
}

type PointDistance struct {
    Point Point
    Cost  int
}

type PointHeap []PointDistance

func (p PointHeap) Len() int {
    return len(p)
}

func (p PointHeap) Less(i, j int) bool {
    return p[i].Cost < p[j].Cost
}

func (p PointHeap) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

func (p *PointHeap) Push(x interface{}) {
    *p = append(*p, x.(PointDistance))
}

func (p *PointHeap) Pop() interface{} {
    old := *p
    n := len(old)
    x := old[n-1]
    *p = old[:n-1]
    return x
}
