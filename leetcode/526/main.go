package leetcode

// https://leetcode.com/problems/beautiful-arrangement/
func countArrangement(n int) int {
    divisibleLists := make([][]int, n+1)
    for i := 1; i <= n; i++ {
        for j := i; j <= n; j += i {
            divisibleLists[i] = append(divisibleLists[i], j)
            if i != j {
                divisibleLists[j] = append(divisibleLists[j], i)
            }
        }
    }
    mask := bitmask(0)
    dp := map[bitmask]int{}
    result := numberOfArrangement(1, n, mask, divisibleLists, dp)
    return result
}

func numberOfArrangement(currentIndex int, lastIndex int, mask bitmask, divisibleLists [][]int, dp map[bitmask]int) int {
    if currentIndex > lastIndex {
        return 1
    }
    if val, ok := dp[mask]; ok {
        return val
    }
    result := 0
    for _, number := range divisibleLists[currentIndex] {
        if !mask.isAvailable(number) {
            continue
        }
        result += numberOfArrangement(currentIndex+1, lastIndex, mask.use(number), divisibleLists, dp)
    }
    dp[mask] = result
    return result
}

type bitmask uint64

func (b bitmask) isAvailable(idx int) bool {
    return (b & (1 << idx)) == 0
}

func (b bitmask) use(idx int) bitmask {
    return b | (1 << idx)
}
