package leetcode

// https://leetcode.com/problems/partition-array-for-maximum-sum/
func maxSumAfterPartitioning(arr []int, maxPartitionMember int) int {
    return maxSumAfterPartitioningHelper(len(arr)-1, arr, maxPartitionMember, memo{})
}

func maxSumAfterPartitioningHelper(idx int, arr []int, maxPartitionMember int, dp memo) int {
    if idx < 0 {
        return 0
    }
    if val, ok := dp[idx]; ok {
        return val
    }
    result := 0
    for i := 1; i <= maxPartitionMember; i++ {
        if (idx - i + 1) < 0 {
            break
        }
        maxInRange := findMaxInArray(arr, idx-i+1, idx)
        res := maxSumAfterPartitioningHelper(idx-i, arr, maxPartitionMember, dp) + maxInRange*i
        result = max(res, result)
    }
    dp[idx] = result
    return result
}

func findMaxInArray(arr []int, left, right int) int {
    res := 0
    for i := left; i <= right; i++ {
        res = max(res, arr[i])
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

type memo map[int]int
