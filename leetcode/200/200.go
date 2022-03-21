package leetcode

import "errors"

const LAND = '1'
const WATER = '0'

// https://leetcode.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	numberOfIslands := 0
	currentMap := newGridMap(grid)
	for row := 0; row < currentMap.GetRowSize(); row++ {
		for column := 0; column < currentMap.GetColumnSize(); column++ {
			currentPoint := newPoint(row, column)
			if currentMap.IsLand(currentPoint) {
				numberOfIslands++
				fillLandWithWater(currentMap, currentPoint)
			}
		}
	}
	return numberOfIslands
}

func fillLandWithWater(currentMap *gridMap, startingPoint point) {
	queue := newQueueOfPoint()
	queue.Enqueue(startingPoint)
	for !queue.IsEmpty() {
		currentPoint := queue.Dequeue()
		if currentMap.IsWater(currentPoint) {
			continue
		}
		currentMap.SetAsWater(currentPoint)
		pointAbove, err := currentMap.GetPointAbove(currentPoint)
		if err == nil {
			queue.Enqueue(pointAbove)
		}
		pointBelow, err := currentMap.GetPointBelow(currentPoint)
		if err == nil {
			queue.Enqueue(pointBelow)
		}
		pointLeft, err := currentMap.GetPointToTheLeft(currentPoint)
		if err == nil {
			queue.Enqueue(pointLeft)
		}
		pointRight, err := currentMap.GetPointToTheRight(currentPoint)
		if err == nil {
			queue.Enqueue(pointRight)
		}
	}
}

// gridMap struct section
type gridMap struct {
	grid [][]byte
}

func (m *gridMap) GetPointAbove(p point) (point, error) {
	if p.row == 0 {
		return point{}, errors.New("no point above")
	}
	return newPoint(p.row-1, p.column), nil
}

func (m *gridMap) GetPointBelow(p point) (point, error) {
	if p.row == m.GetRowSize()-1 {
		return point{}, errors.New("no point below")
	}
	return newPoint(p.row+1, p.column), nil
}

func (m *gridMap) GetPointToTheLeft(p point) (point, error) {
	if p.column == 0 {
		return point{}, errors.New("no point to the left")
	}
	return newPoint(p.row, p.column-1), nil
}

func (m *gridMap) GetPointToTheRight(p point) (point, error) {
	if p.column == m.GetColumnSize()-1 {
		return point{}, errors.New("no point to the right")
	}
	return newPoint(p.row, p.column+1), nil
}

func (m *gridMap) GetRowSize() int {
	return len(m.grid)
}

func (m *gridMap) GetColumnSize() int {
	return len(m.grid[0])
}

func (m *gridMap) IsLand(point point) bool {
	return m.grid[point.row][point.column] == LAND
}

func (m *gridMap) IsWater(point point) bool {
	return m.grid[point.row][point.column] == WATER
}

func (m *gridMap) SetAsWater(point point) {
	m.grid[point.row][point.column] = WATER
}

func newGridMap(grid [][]byte) *gridMap {
	return &gridMap{grid: grid}
}

// point struct section
type point struct {
	row, column int
}

func newPoint(row, column int) point {
	return point{
		row:    row,
		column: column,
	}
}

// queue struct section
type queueOfPoint struct {
	container []point
}

func (q *queueOfPoint) Enqueue(point point) {
	q.container = append(q.container, point)
}

func (q *queueOfPoint) Dequeue() point {
	firstPoint := q.container[0]
	q.container = q.container[1:]
	return firstPoint
}

func (q *queueOfPoint) IsEmpty() bool {
	return len(q.container) == 0
}

func newQueueOfPoint() *queueOfPoint {
	return &queueOfPoint{}
}
