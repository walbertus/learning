package leetcode
//
//// https://leetcode.com/problems/longest-palindromic-substring
//func longestPalindrome(s string) string {
//	longestLength := 1
//	longestLeftIndex := 0
//	longestRightIndext := 0
//	checker := newPalindromeChecker()
//	for i := 0; i < len(s); i++ {
//		for j := i + 1; j < len(s); j++ {
//			if longestLength >= (j-i+1) {
//				continue
//			}
//			if checker.isSubstringPalindrome(s, i, j) && longestLength < (j-i+1) {
//				longestLength = j-i+1
//				longestLeftIndex = i
//				longestRightIndext = j
//			}
//		}
//	}
//	return s[longestLeftIndex:longestRightIndext+1]
//}
//
//type palindromeChecker struct {
//}
//
//func (c *palindromeChecker) isSubstringPalindrome(str string, leftIndex, rightIndex int) bool {
//	if leftIndex == rightIndex {
//		return true
//	}
//	if str[leftIndex] != str[rightIndex] {
//		return false
//	}
//	if rightIndex-leftIndex == 1 {
//		return true
//	}
//	return c.isSubstringPalindrome(str, leftIndex+1, rightIndex-1)
//}
//
//func newPalindromeChecker() *palindromeChecker {
//	return &palindromeChecker{}
//}
