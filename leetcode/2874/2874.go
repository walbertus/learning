package leetcode

// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-ii/
func maximumTripletValue(nums []int) int64 {
	maxValueFromBeginning := make([]int, len(nums))
	maxValueFromBeginning[0] = nums[0]

	maxValueToEnd := make([]int, len(nums))
	maxValueToEnd[len(nums)-1] = nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		maxValueFromBeginning[i] = maxInt(maxValueFromBeginning[i-1], nums[i])

		currentIdx := len(nums) - 1 - i
		maxValueToEnd[currentIdx] = maxInt(maxValueToEnd[currentIdx+1], nums[currentIdx])
	}
	result := int64(0)
	for i := 1; i < len(nums)-1; i++ {
		currentResult := int64((maxValueFromBeginning[i-1] - nums[i])) * int64(maxValueToEnd[i+1])
		result = maxInt64(result, currentResult)
	}

	return result
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
