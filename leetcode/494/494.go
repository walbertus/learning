package leetcode

// https://leetcode.com/problems/target-sum/
func findTargetSumWays(nums []int, target int) int {
	memo := newMemo()
	memo.incrementInKey(0, nums[0], 1)
	memo.incrementInKey(0, -nums[0], 1)
	for i := 1; i < len(nums); i++ {
		resultInPreviousLocation := memo.getInLocation(i - 1)
		for val, totalCombination := range resultInPreviousLocation {
			memo.incrementInKey(i, val-nums[i], totalCombination)
			memo.incrementInKey(i, val+nums[i], totalCombination)
		}
		memo.deleteInLocation(i - 1)
	}
	return memo.getInKey(len(nums)-1, target)
}

type memo struct {
	container map[int]map[int]int
}

func (m *memo) incrementInKey(location, value, increment int) {
	if _, ok := m.container[location]; !ok {
		m.container[location] = map[int]int{}
	}
	m.container[location][value] += increment
}

func (m *memo) getInLocation(location int) map[int]int {
	if _, ok := m.container[location]; !ok {
		m.container[location] = map[int]int{}
	}
	return m.container[location]
}

func (m *memo) deleteInLocation(location int) {
	delete(m.container, location)
}

func (m *memo) getInKey(location, value int) int {
	if _, ok := m.container[location]; !ok {
		return 0
	}
	return m.container[location][value]
}

func newMemo() *memo {
	return &memo{
		container: map[int]map[int]int{},
	}
}
