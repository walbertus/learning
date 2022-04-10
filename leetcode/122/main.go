package leetcode

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
    result := 0
    currentMin := prices[0]
    currentMax := prices[0]
    for i := 0; i < len(prices); i++ {
        price := prices[i]
        if price >= currentMax {
            currentMax = price
        } else {
            result += currentMax - currentMin
            if i != len(prices)-1 {
                currentMax = prices[i+1]
                currentMin = prices[i+1]
            } else {
                currentMax, currentMin = 0, 0
            }
        }
        currentMin = min(currentMin, price)
    }
    return result + currentMax - currentMin
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
