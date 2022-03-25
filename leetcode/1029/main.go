package leetcode

import "sort"

// https://leetcode.com/problems/two-city-scheduling/
func twoCitySchedCost(costs [][]int) int {
    sort.Sort(CostsByDiff(costs))
    result := 0
    for i := 0; i < len(costs)/2; i++ {
        result += costs[i][1] + costs[len(costs)-i-1][0]
    }
    return result
}

type CostsByDiff [][]int

func (c CostsByDiff) Len() int {
    return len(c)
}

func (c CostsByDiff) Less(i, j int) bool {
    diffI := c[i][1] - c[i][0]
    diffJ := c[j][1] - c[j][0]
    return diffI < diffJ
}

func (c CostsByDiff) Swap(i, j int) {
    c[i], c[j] = c[j], c[i]
}
