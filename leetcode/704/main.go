package leetcode

// https://leetcode.com/problems/binary-search/
func search(nums []int, target int) int {
    leftPointer := 0
    rightPointer := len(nums)-1
    for leftPointer <= rightPointer {
        midPoint := (leftPointer + rightPointer)/ 2
        if nums[midPoint] == target{
            return midPoint
        }
        if nums[midPoint] < target {
            leftPointer = midPoint + 1
        } else {
            rightPointer = midPoint - 1
        }
    }
    return -1
}
