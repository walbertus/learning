package leetcode

// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	counterNumbers := map[int]int{}
	maxOccurence := 0
	for _, val := range nums {
		counterNumbers[val]++
		maxOccurence = max(maxOccurence, counterNumbers[val])
	}
	occurenceID := make([][]int, maxOccurence + 1)
	for id, val := range counterNumbers {
		occurenceID[val] = append(occurenceID[val], id)
	}
	var result []int
	idx := maxOccurence
	for len(result) < k {
		if len(occurenceID[idx]) != 0 {
			result = append(result, occurenceID[idx]...)
		}
		idx--
	}
	return result
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
