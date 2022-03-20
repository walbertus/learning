package leetcode

import "fmt"

// https://leetcode.com/problems/longest-palindromic-substring
func longestPalindrome(s string) string {
	longestLength := 1
	longestLeftIndex := 0
	longestRightIndex := 0
	checker := newPalindromeChecker()
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i+longestLength-1; j-- {
			if (j - i + 1) < longestLength {
				continue
			}
			if checker.isSubstringPalindrome(s, i, j) && longestLength < (j-i+1) {
				longestLength = j - i + 1
				longestLeftIndex = i
				longestRightIndex = j
			}
		}
	}
	return s[longestLeftIndex : longestRightIndex+1]
}

type palindromeChecker struct {
	dp map[int]map[int]bool
}

func (c *palindromeChecker) isSubstringPalindrome(str string, leftIndex, rightIndex int) bool {
	if leftIndex == rightIndex {
		return true
	}
	if c.isInDP(leftIndex, rightIndex) {
		return c.dp[leftIndex][rightIndex]
	}
	if str[leftIndex] != str[rightIndex] {
		c.dp[leftIndex][rightIndex] = false
		return false
	}
	if rightIndex-leftIndex == 1 {
		c.dp[leftIndex][rightIndex] = true
		return c.dp[leftIndex][rightIndex]
	}
	c.dp[leftIndex][rightIndex] = c.isSubstringPalindrome(str, leftIndex+1, rightIndex-1)
	return c.dp[leftIndex][rightIndex]
}

func (c *palindromeChecker) isInDP(leftIndex, rightIndex int) bool {
	if _, ok := c.dp[leftIndex]; !ok {
		c.dp[leftIndex] = map[int]bool{}
	}
	_, ok := c.dp[leftIndex][rightIndex]
	return ok
}

func (c *palindromeChecker) generateDPKey(leftIndex, rightIndex int) string {
	return fmt.Sprintf("%d-%d", leftIndex, rightIndex)
}

func newPalindromeChecker() *palindromeChecker {
	return &palindromeChecker{
		dp: map[int]map[int]bool{},
	}
}
