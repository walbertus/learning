package leetcode

// https://leetcode.com/problems/broken-calculator/
func brokenCalc(startValue int, target int) int {
	numberOfOperations := 0
	numberOfTimesTwo := 0
	currentValue := startValue
	for currentValue < target {
		numberOfTimesTwo++
		currentValue *= 2
	}
	numberOfOperations = numberOfTimesTwo
	for currentValue != target {
		currentBiggestTwoSquare := square(2, numberOfTimesTwo)
		for currentValue-currentBiggestTwoSquare >= target {
			currentValue -= currentBiggestTwoSquare
			numberOfOperations++
		}
		if currentValue < target {
			panic("too far")
		}
		numberOfTimesTwo--
	}
	return numberOfOperations
}

func square(base, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= base
	}
	return result
}
