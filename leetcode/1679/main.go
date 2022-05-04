package leetcode

// https://leetcode.com/problems/max-number-of-k-sum-pairs/
func maxOperations(nums []int, k int) int {
    counter := map[int]int{}
    for _, num := range nums {
        counter[num]++
    }
    result := 0
    for key, value := range counter {
        if key >= k || value == 0 {
            continue
        }
        complement, ok := counter[k-key]
        if !ok {
            continue
        }
        if key == k-key {
            currentResult := value / 2
            result += currentResult
            counter[key] -= currentResult
            continue
        }
        currentResult := min(value, complement)
        result += currentResult
        counter[key] -= currentResult
        counter[k-key] -= currentResult
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
