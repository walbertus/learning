package leetcode

// https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/
func minDominoRotations(tops []int, bottoms []int) int {
	result := len(tops)
	found := false
	resultEqualizingTop, ok := numberOfChangesToEqualize(tops, bottoms, tops[0])
	found = found || ok
	if ok {
		resultEqualizingTop = min(resultEqualizingTop, len(tops) - resultEqualizingTop)
		result = min(result, resultEqualizingTop)
	}
	resultEqualizingTopWithBottom, ok := numberOfChangesToEqualize(tops, bottoms, bottoms[0])
	found = found || ok
	if ok {
		resultEqualizingTopWithBottom = min(resultEqualizingTopWithBottom, len(tops) - resultEqualizingTopWithBottom)
		result = min(result, resultEqualizingTopWithBottom)
	}
	resultEqualizingBottom, ok := numberOfChangesToEqualize(bottoms, tops, bottoms[0])
	found = found || ok
	if ok {
		resultEqualizingBottom = min(resultEqualizingBottom, len(tops) - resultEqualizingBottom)
		result = min(result, resultEqualizingBottom)
	}
	resultEqualizingBottomWithTop, ok := numberOfChangesToEqualize(bottoms, tops, tops[0])
	found = found || ok
	if ok {
		resultEqualizingBottomWithTop = min(resultEqualizingBottomWithTop, len(tops) - resultEqualizingBottomWithTop)
		result = min(result, resultEqualizingBottomWithTop)
	}
	if found {
		return result
	} else {
		return -1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numberOfChangesToEqualize(source, destination []int, target int) (int, bool) {
	numberOfRotation := 0
	for i := 0; i < len(source); i++ {
		if source[i] == target {
			continue
		}
		if destination[i] == target {
			numberOfRotation++
		} else {
			return -1, false
		}
	}
	return numberOfRotation, true
}
