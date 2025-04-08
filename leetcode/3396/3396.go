package leetcode

func minimumOperations(nums []int) int {
	counter := make([]int, 101)
	for i := len(nums) - 1; i >= 0; i-- {
		counter[nums[i]]++
		if counter[nums[i]] > 1 {
			return (i/3) + 1
		}
	}
	return 0
}
