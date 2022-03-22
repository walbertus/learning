package leetcode

import (
	"fmt"
)

// https://leetcode.com/problems/ones-and-zeroes/
func findMaxForm(strs []string, m int, n int) int {
	var binariesMetadata []binaryStringMetadata
	for _, str := range strs {
		binariesMetadata = append(binariesMetadata, newBinaryStringMetadata(str))
	}
	subsetCounter := newBinarySubsetCounter(binariesMetadata)
	return subsetCounter.countSubnet(0, m, n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type knapsackBinarySubset struct {
	memo             map[int]map[int]map[int]int
	binariesMetadata []binaryStringMetadata
}

func (s *knapsackBinarySubset) countSubnet(
	currentPosition int,
	maxNumberOfZero int,
	maxNumberOfOne int,
) int {
	if currentPosition == len(s.binariesMetadata) {
		return 0
	}
	result, ok := s.getValueFromMemo(currentPosition, maxNumberOfZero, maxNumberOfOne)
	if ok {
		return result
	}
	result = s.countSubnet(currentPosition+1, maxNumberOfZero, maxNumberOfOne)
	if !s.isOverflowIfTaken(currentPosition, maxNumberOfZero, maxNumberOfOne) {
		maxNumberOfZeroLeft := maxNumberOfZero - s.binariesMetadata[currentPosition].numberOfZero
		maxNumberOfOneLeft := maxNumberOfOne - s.binariesMetadata[currentPosition].numberOfOne
		resultIfTaken := s.countSubnet(currentPosition+1, maxNumberOfZeroLeft, maxNumberOfOneLeft)
		result = max(result, resultIfTaken+1)
	}
	s.setValueToMemo(currentPosition, maxNumberOfZero, maxNumberOfOne, result)
	return result
}

func (s *knapsackBinarySubset) isOverflowIfTaken(currentPosition, maxNumberOfZero, maxNumberOfOne int) bool {
	return ((maxNumberOfZero - s.binariesMetadata[currentPosition].numberOfZero) < 0) ||
		((maxNumberOfOne - s.binariesMetadata[currentPosition].numberOfOne) < 0)

}

func (s *knapsackBinarySubset) setValueToMemo(
	currentPosition,
	maxNumberOfZero,
	maxNumberOfOne,
	value int,
) {
	if _, ok := s.memo[currentPosition]; !ok {
		s.memo[currentPosition] = map[int]map[int]int{}
	}
	if _, ok := s.memo[currentPosition][maxNumberOfZero]; !ok {
		s.memo[currentPosition][maxNumberOfZero] = map[int]int{}
	}
	s.memo[currentPosition][maxNumberOfZero][maxNumberOfOne] = value
}

func (s *knapsackBinarySubset) getValueFromMemo(
	currentPosition,
	maxNumberOfZero,
	maxNumberOfOne int,
) (int, bool) {
	if _, ok := s.memo[currentPosition]; !ok {
		return 0, false
	}
	if _, ok := s.memo[currentPosition][maxNumberOfZero]; !ok {
		return 0, false
	}
	if val, ok := s.memo[currentPosition][maxNumberOfZero][maxNumberOfOne]; !ok {
		return 0, false
	} else {
		return val, true
	}
}

func newBinarySubsetCounter(binariesMetadata []binaryStringMetadata) *knapsackBinarySubset {
	return &knapsackBinarySubset{
		memo:             map[int]map[int]map[int]int{},
		binariesMetadata: binariesMetadata,
	}
}

type binaryStringMetadata struct {
	numberOfZero int
	numberOfOne  int
}

func newBinaryStringMetadata(binaryString string) binaryStringMetadata {
	numberOfOne := 0
	numberOfZero := 0
	for _, char := range binaryString {
		if char == '1' {
			numberOfOne++
		} else if char == '0' {
			numberOfZero++
		} else {
			panic(fmt.Sprintf("unknown rune %c", char))
		}
	}
	return binaryStringMetadata{
		numberOfZero: numberOfZero,
		numberOfOne:  numberOfOne,
	}
}
