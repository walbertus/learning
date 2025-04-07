package leetcode

// https://leetcode.com/problems/partition-equal-subset-sum
func canPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}
	if total&1 == 1 {
		return false
	}
	return solution(nums, 0, total>>1, newMemo(len(nums)))
}

func solution(nums []int, index int, left int, dp *memo) bool {
	if left == 0 {
		return true
	}
	if index >= len(nums) || left < 0 {
		return false
	}
	if v, ok := dp.get(index, left); ok {
		return v
	}
	var result bool
	if solution(nums, index+1, left-nums[index], dp) {
		result = true
	} else {
		result = solution(nums, index+1, left, dp)
	}
	dp.set(index, left, result)
	return result
}

type memo struct {
	container []map[int]bool
}

func newMemo(size int) *memo {
	m := &memo{
		container: make([]map[int]bool, size),
	}
	for i := range m.container {
		m.container[i] = make(map[int]bool)
	}
	return m
}

func (m *memo) get(index int, left int) (bool, bool) {
	if index >= len(m.container) {
		return false, false
	}
	v, ok := m.container[index][left]
	return v, ok
}

func (m *memo) set(index int, left int, value bool) {
	if index >= len(m.container) {
		return
	}
	m.container[index][left] = value
}
