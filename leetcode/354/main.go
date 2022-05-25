package leetcode

import "sort"

// https://leetcode.com/problems/russian-doll-envelopes
func maxEnvelopes(envelopes [][]int) int {
    sort.Sort(CustomSort(envelopes))
    var nums []int
    for _, envelope := range envelopes {
        nums = append(nums, envelope[1])
    }
    return LIS(nums)
}

func CeilIndex(tails []int, num int) int {
    l, r := 0, len(tails)-1
    res := 0
    for l <= r {
        mid := (l + r) / 2
        if tails[mid] < num {
            l = mid + 1
        } else {
            res = mid
            r = mid - 1
        }
    }
    return res
}

func LIS(nums []int) int {
    tails := []int{nums[0]}
    for _, num := range nums {
        if num < tails[0] {
            tails[0] = num
            continue
        }
        if num > tails[len(tails)-1] {
            tails = append(tails, num)
            continue
        }
        insertIndex := CeilIndex(tails, num)
        if tails[insertIndex] > num {
            tails[insertIndex] = num
        }
    }
    return len(tails)
}

type CustomSort [][]int

func (c CustomSort) Len() int {
    return len(c)
}

func (c CustomSort) Swap(i, j int) {
    c[i], c[j] = c[j], c[i]
}

func (c CustomSort) Less(i, j int) bool {
    if c[i][0] < c[j][0] {
        return true
    }
    if c[i][0] > c[j][0] {
        return false
    }
    return c[i][1] > c[j][1]
}
