package leetcode

import "sort"

func checkValidCuts(n int, rectangles [][]int) bool {
	xCounter := make(map[int]int)
	yCounter := make(map[int]int)

	for _, rectangle := range rectangles {
		x1, y1, x2, y2 := rectangle[0], rectangle[1], rectangle[2], rectangle[3]

		xCounter[x1+1]++
		xCounter[x2]--

		yCounter[y1+1]++
		yCounter[y2]--
	}

	if lineSweep(xCounter) >= 2 {
		return true
	}

	return lineSweep(yCounter) >= 2
}

func lineSweep(counter map[int]int) int {
	validCut := 0
	count := 0

	numberToVisit := make([]int, 0, len(counter))
	for key, _ := range counter {
		numberToVisit = append(numberToVisit, key)
	}

	sort.Ints(numberToVisit)

	for i := 0; i < len(numberToVisit)-1; i++ {
		now := numberToVisit[i]
		count += counter[now]
		if count == 0 {
			validCut++
		}
	}
	return validCut
}
