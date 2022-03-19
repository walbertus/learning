package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	charSet := newSet()

	longestLength := 1
	currentLength := 1
	leftPointerIndex := 0
	rightPointerIndex := 0
	for {
		if leftPointerIndex == len(s) {
			break
		}
		isLastChar := rightPointerIndex == len(s)
		if isLastChar || charSet.Exist(rune(s[rightPointerIndex])) {
			leftmostChar := rune(s[leftPointerIndex])
			charSet.Delete(leftmostChar)
			leftPointerIndex++
			currentLength--
			longestLength = max(longestLength, currentLength)
		} else {
			longestLength = max(longestLength, currentLength)
			charSet.Add(rune(s[rightPointerIndex]))
			rightPointerIndex++
			currentLength++
		}
	}
	longestLength = max(longestLength, currentLength)
	return longestLength
}

type set struct {
	container map[rune]bool
}

func (s *set) Exist(key rune) bool {
	return s.container[key]
}

func (s *set) Add(key rune) {
	s.container[key] = true
}

func (s *set) Delete(key rune) {
	s.container[key] = false
}

func newSet() *set {
	return &set{container: map[rune]bool{}}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
