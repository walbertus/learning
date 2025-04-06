package leetcode

import "sort"

// TODO: this got accepted but very slow, I don't think this is the expected solution
// https://leetcode.com/problems/largest-divisible-subset
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	return solution(nums, 1, 0, newMemo(len(nums)))
}

func solution(nums []int, timesValue int, idx int, dp *memo) []int {
	if idx == len(nums) {
		return []int{}
	}
	if value, ok := dp.get(idx, timesValue); ok {
		return value
	}
	without := solution(nums, timesValue, idx+1, dp)
	result := without
	if nums[idx]%timesValue == 0 {
		with := solution(nums, timesValue*(nums[idx]/timesValue), idx+1, dp)
		with = append([]int{nums[idx]}, with...)
		if len(with) >= len(without) {
			result = with
		}
	}
	dp.set(idx, timesValue, result)
	return result
}

type memo struct {
	container []map[int][]int
}

func newMemo(n int) *memo {
	m := make([]map[int][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make(map[int][]int)
	}
	return &memo{container: m}
}

func (m *memo) set(idx, timesValue int, value []int) {
	if idx >= len(m.container) {
		panic("not expected")
	}
	m.container[idx][timesValue] = value
}

func (m *memo) get(idx, timesValue int) ([]int, bool) {
	if idx >= len(m.container) {
		panic("not expected")
	}
	value, ok := m.container[idx][timesValue]
	return value, ok
}
