package leetcode

func maxArea(heights []int) int {
    start := 0
    end := len(heights) - 1
    area := 0
    for start < end {
        area = max(area, (end-start)* min(heights[end], heights[start]))
        if heights[end] < heights[start] {
            end--
        } else {
            start++
        }
    }
    return area
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
