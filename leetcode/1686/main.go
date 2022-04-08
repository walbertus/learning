package leetcode

import (
    "sort"
)

type diffData struct {
    diff, index int
}

type byDiff []diffData

func (b byDiff) Len() int {
    return len(b)
}

func (b byDiff) Less(i, j int) bool {
    return b[i].diff > b[j].diff
}

func (b byDiff) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}

// https://leetcode.com/problems/stone-game-vi/
func stoneGameVI(aliceValues []int, bobValues []int) int {
    valuesLength := len(aliceValues)
    diffValues := make([]diffData, valuesLength)
    for i := 0; i < valuesLength; i++ {
        diffValues[i] = diffData{
            diff:  intAbs(aliceValues[i] + bobValues[i]),
            index: i,
        }
    }
    total := 0
    sort.Sort(byDiff(diffValues))
    for i := 0; i < valuesLength/2; i++ {
        aliceIndex := diffValues[2*i].index
        bobIndex := diffValues[2*i+1].index
        total += aliceValues[aliceIndex] - bobValues[bobIndex]
    }
    if valuesLength%2 != 0 {
        total += aliceValues[diffValues[valuesLength-1].index]
    }
    if total == 0 {
        return 0
    }
    if total < 0 {
        return -1
    }
    return 1
}

func intAbs(value int) int {
    if value < 0 {
        return -value
    }
    return value
}
