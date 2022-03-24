package leetcode

import "math"

func getMoneyAmount(numberLimit int) int {
	counter := newCounter()
	return counter.countTotalCost(1, numberLimit)
}

type counter struct {
	memo *memo
}

func (c *counter) countTotalCost(lowestNumber, highestNumber int) int {
	if lowestNumber == highestNumber {
		return 0
	}
	if value, ok := c.memo.getValue(lowestNumber, highestNumber); ok {
		return value
	}
	result := math.MaxInt
	for i := lowestNumber; i <= highestNumber; i++ {
		costWhenLeft := -1
		if i != lowestNumber {
			costWhenLeft = c.countTotalCost(lowestNumber, i-1) + i
		}
		costWhenRight := -1
		if i != highestNumber {
			costWhenRight = c.countTotalCost(i+1, highestNumber) + i
		}
		result = min(result, max(costWhenLeft, costWhenRight))
	}
	c.memo.setValue(lowestNumber, highestNumber, result)
	return result
}

func newCounter() *counter {
	return &counter{memo: newMemo(201)}
}

type memo struct {
	containerArray [][]int
}

func (m *memo) setValue(key1, key2, value int) {
	m.containerArray[key1][key2] = value
}

func (m *memo) getValue(key1, key2 int) (int, bool) {
	if m.containerArray[key1][key2] == 0 {
		return 0, false
	}
	return m.containerArray[key1][key2], true
}

func newMemo(size int) *memo {
	containerArray := make([][]int, size+1)
	for i := range containerArray {
		containerArray[i] = make([]int, size+1)
		for j := range containerArray[i] {
			containerArray[i][j] = 0
		}
	}
	return &memo{
		containerArray: containerArray,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
