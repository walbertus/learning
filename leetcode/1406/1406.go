package leetcode

import (
	"math"
)

// https://leetcode.com/problems/stone-game-iii
func stoneGameIII(stoneValue []int) string {
	dp := newMemo(len(stoneValue))

	aliceScore := maxStone(stoneValue, 0, dp)
	if aliceScore > 0 {
		return "Alice"
	} else if aliceScore < 0 {
		return "Bob"
	}
	return "Tie"
}

func maxStone(stoneValue []int, idx int, dp *memo) int {
	if ok, value := dp.get(idx); ok {
		return value
	}

	result := math.MinInt
	stoneTake := 0
	for i := 1; i <= 3; i++ {
		if idx+i-1 >= len(stoneValue) {
			break
		}
		stoneTake += stoneValue[idx+i-1]
		currentValue := stoneTake
		if idx+i < len(stoneValue) {
			currentValue -= maxStone(stoneValue, idx+i, dp)
		}
		result = maxInt(result, currentValue)
	}

	dp.set(idx, result)
	return result
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type memo struct {
	value []int
}

func newMemo(size int) *memo {
	value := make([]int, size)
	for i := 0; i < size; i++ {
		value[i] = math.MinInt
	}
	return &memo{
		value: value,
	}
}

func (m *memo) get(idx int) (bool, int) {
	if m.value[idx] == math.MinInt {
		return false, 0
	}
	return true, m.value[idx]
}

func (m *memo) set(idx int, value int) {
	m.value[idx] = value
}
