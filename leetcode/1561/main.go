package leetcode

import "sort"

// https://leetcode.com/problems/maximum-number-of-coins-you-can-get/
func maxCoins(piles []int) int {
    sort.Ints(piles)
    res := 0
    pilesLength := len(piles)
    numberOfPileTaken := pilesLength / 3
    pileToTake := pilesLength-2
    for i := 0; i < numberOfPileTaken; i++ {
        res += piles[pileToTake]
        pileToTake -= 2
    }
    return res
}
