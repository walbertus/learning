package leetcode

const minimumPossibleHealth = -900000000

// https://leetcode.com/problems/dungeon-game/
func calculateMinimumHP(dungeonData [][]int) int {
	dungeon := newDungeon(dungeonData, true)
	startPoint := newPoint(0, 0)
	endPoint := newPoint(len(dungeonData)-1, len(dungeonData[0])-1)
	result := dungeon.calculateMinimumHPSpend(startPoint, endPoint)
	return (-result) + 1
}

type point struct {
	row, column int
}

func (p point) equal(otherPoint point) bool {
	if p.row != otherPoint.row {
		return false
	}
	return p.column == otherPoint.column
}

func newPoint(row, column int) point {
	return point{
		row:    row,
		column: column,
	}
}

type dungeon struct {
	dungeonData   [][]int
	dungeonRow    int
	dungeonColumn int
	memo          *memo
	enableMemo    bool
}

func (d *dungeon) calculateMinimumHPSpend(startPoint, endPoint point) int {
	damageAtThisPoint := d.damageInPoint(startPoint)
	if startPoint.equal(endPoint) {
		return min(damageAtThisPoint, 0)
	}
	if d.enableMemo {
		if val, ok := d.memo.getValue(startPoint); ok {
			return val
		}
	}
	belowPoint := newPoint(startPoint.row+1, startPoint.column)
	belowHP := minimumPossibleHealth
	if d.isInDungeon(belowPoint) {
		belowHP = d.calculateMinimumHPSpend(belowPoint, endPoint) + damageAtThisPoint
	}
	rightPoint := newPoint(startPoint.row, startPoint.column+1)
	rightHP := minimumPossibleHealth
	if d.isInDungeon(rightPoint) {
		rightHP = d.calculateMinimumHPSpend(rightPoint, endPoint) + damageAtThisPoint
	}
	minimumHPNeeded := min(0, max(belowHP, rightHP))
	if d.enableMemo {
		d.memo.setValue(startPoint, minimumHPNeeded)
	}
	return minimumHPNeeded
}

func (d *dungeon) damageInPoint(p point) int {
	return d.dungeonData[p.row][p.column]
}

func (d *dungeon) isInDungeon(p point) bool {
	if p.row < 0 || p.column < 0 {
		return false
	}
	return p.row < d.dungeonRow && p.column < d.dungeonColumn
}

func newDungeon(dungeonData [][]int, enableMemo bool) *dungeon {
	return &dungeon{
		dungeonData:   dungeonData,
		dungeonRow:    len(dungeonData),
		dungeonColumn: len(dungeonData[0]),
		memo:          newMemo(),
		enableMemo:    enableMemo,
	}
}

type memo struct {
	container map[int]map[int]int
}

func (m *memo) setValue(key point, value int) {
	if _, ok := m.container[key.row]; !ok {
		m.container[key.row] = map[int]int{}
	}
	m.container[key.row][key.column] = value
}

func (m *memo) getValue(key point) (int, bool) {
	if _, ok := m.container[key.row]; !ok {
		return 0, false
	}
	if _, ok := m.container[key.row][key.column]; !ok {
		return 0, false
	}
	return m.container[key.row][key.column], true
}

func newMemo() *memo {
	return &memo{
		container: map[int]map[int]int{},
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
