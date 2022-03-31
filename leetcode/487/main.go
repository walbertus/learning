package leetcode

// https://leetcode.com/problems/predict-the-winner/
func PredictTheWinner(nums []int) bool {
    answer := make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        answer[i] = make([]int, len(nums))
        answer[i][i] = nums[i]
    }

    sum := make(totalInRange, len(nums))
    sum.preComputeTotal(nums)
    for length := 1; length < len(nums); length++ {
        for i := 0; i < len(nums)-length; i++ {
            totalStoneLeft := sum.getInRange(i, i+length)
            takeLeft := totalStoneLeft - answer[i+1][i+length]
            takeRight := totalStoneLeft - answer[i][i+length-1]
            answer[i][i+length] = max(takeLeft, takeRight)
        }
    }
    totalStone := sum.getInRange(0, len(nums)-1)
    maxPossible := answer[0][len(nums)-1]
    return maxPossible >= (totalStone+1)/2
}

type totalInRange [][]int

func (t totalInRange) preComputeTotal(nums []int) {
    for i := 0; i < len(nums); i++ {
        t[i] = make([]int, len(nums))
        t[i][i] = nums[i]
        for j := i + 1; j < len(nums); j++ {
            t[i][j] = nums[j] + t[i][j-1]
        }
    }
}

func (t totalInRange) getInRange(start int, end int) int {
    return t[start][end]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
