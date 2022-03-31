package leetcode

import "sort"

// https://leetcode.com/problems/delete-and-earn/submissions/
func deleteAndEarn(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    numberOccurrence := occurrenceData{}
    var uniqueNums []int
    for _, val := range nums {
        if numberOccurrence[val] == 0 {
            uniqueNums = append(uniqueNums, val)
        }
        numberOccurrence.add(val)
    }
    sort.Ints(uniqueNums)
    return dp(0, uniqueNums, numberOccurrence, memo{})
}

type memo map[int]int

func dp(id int, uniqueNums []int, numberOccurrence occurrenceData, memo memo) int {
    if id >= len(uniqueNums) {
        return 0
    }
    if value, ok := memo[id]; ok {
        return value
    }

    currentNum := uniqueNums[id]
    if id == len(uniqueNums)-1 {
        result := currentNum * numberOccurrence[currentNum]
        memo[id] = result
        return result
    }
    whenTaken := 0
    if uniqueNums[id+1]-currentNum == 1 {
        whenTaken = currentNum * numberOccurrence[currentNum]
        whenTaken += dp(id+2, uniqueNums, numberOccurrence, memo)
    } else {
        whenTaken = currentNum * numberOccurrence[currentNum]
        whenTaken += dp(id+1, uniqueNums, numberOccurrence, memo)
    }
    whenNotTaken := dp(id+1, uniqueNums, numberOccurrence, memo)
    result := max(whenTaken, whenNotTaken)
    memo[id] = result
    return result
}

type occurrenceData map[int]int

func (o occurrenceData) add(num int) {
    o[num]++
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
