package leetcode

import (
    "math"
    "sort"
)

// https://leetcode.com/problems/next-permutation/
func nextPermutation(nums []int) {
    if len(nums) == 1 {
        return
    }
    if isSortedDescending(nums) {
        sort.Ints(nums)
        return
    }
    idxToIncrement := -1
    for idx := len(nums) - 2; idx >= 0; idx-- {
        if nums[idx] < nums[idx+1] {
            idxToIncrement = idx
            break
        }
    }
    idxToSwap := -1
    currentMinimum := math.MaxInt
    for idx := idxToIncrement+1; idx < len(nums); idx++ {
        if nums[idx] <= currentMinimum && nums[idx] > nums[idxToIncrement] {
            currentMinimum = nums[idx]
            idxToSwap = idx
        }
    }
    swap(nums, idxToSwap, idxToIncrement)
    for i := 1; i < (len(nums)-idxToIncrement+1)/2; i++ {
        swap(nums, idxToIncrement+i, len(nums)-i)
    }
}

func swap(nums []int, a, b int) {
    nums[a], nums[b] = nums[b], nums[a]
}

func isSortedDescending(nums []int) bool {
    return sort.SliceIsSorted(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })
}
