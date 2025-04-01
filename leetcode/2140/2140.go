package leetcode

func mostPoints(questions [][]int) int64 {
	maxPoints := make([]int64, len(questions))
	maxPoints[len(questions)-1] = int64(questions[len(questions)-1][0])

	for i := len(questions) - 2; i >= 0; i-- {
		pointsIfSkipped := int64(maxPoints[i+1])
		pointsIfNotSkipped := int64(questions[i][0])

		next := i + questions[i][1] + 1
		if next < len(questions) {
			pointsIfNotSkipped += maxPoints[next]
		}

		maxPoints[i] = maxInt64(pointsIfSkipped, pointsIfNotSkipped)
	}
	return maxPoints[0]
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
