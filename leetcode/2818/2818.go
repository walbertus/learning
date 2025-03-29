package leetcode

import (
	"math"
	"sort"
)

const MOD = 1000000007

func maximumScore(nums []int, k int) int {
	primeScores := generatePrimeScores(nums)
	leftReaches := generateLeftReaches(primeScores)
	rightReaches := generateRightReaches(primeScores)

	// data structure to sort the biggest number and keep the index
	// index 0 is the number, index 1 is the real index
	biggestNumberMarker := make([][]int, len(primeScores))
	for i := 0; i < len(primeScores); i++ {
		biggestNumberMarker[i] = make([]int, 2)
		biggestNumberMarker[i][0] = nums[i]
		biggestNumberMarker[i][1] = i
	}

	// descending sort
	sort.Slice(biggestNumberMarker, func(i, j int) bool {
		return biggestNumberMarker[i][0] > biggestNumberMarker[j][0]
	})

	finalScore := int64(1)
	i := 0
	for k > 0 {
		number := biggestNumberMarker[i][0]
		idx := biggestNumberMarker[i][1]
		leftLength := idx - leftReaches[idx] - 1
		rightLength := rightReaches[idx] - idx - 1

		totalCombination := leftLength + rightLength + (leftLength * rightLength) + 1
		totalCombination = int(math.Min(float64(totalCombination), float64(k)))

		finalScore = multiply(finalScore, power(int64(number), int64(totalCombination)))
		k -= totalCombination
		i++
	}

	return int(finalScore)
}

func multiply(a, b int64) int64 {
	result := a * b % MOD
	return result
}

func power(base, exp int64) int64 {
	if exp == 1 {
		return base
	}
	result := multiply(base, base)
	result = power(result, exp>>1)
	if exp&1 == 1 {
		result = multiply(result, base)
		exp--
	}
	return result
}

func generateLeftReaches(primeScores []int) []int {
	leftReaches := make([]int, len(primeScores))
	leftReaches[0] = -1
	for i := 1; i < len(primeScores); i++ {
		if primeScores[i-1] >= primeScores[i] {
			leftReaches[i] = i - 1
			continue
		}

		j := i - 1
		for j >= 0 && primeScores[j] < primeScores[i] {
			j = leftReaches[j]
		}
		leftReaches[i] = j
	}
	return leftReaches
}

func generateRightReaches(primeScores []int) []int {
	rightReaches := make([]int, len(primeScores))
	rightReaches[len(primeScores)-1] = len(primeScores)
	for i := len(primeScores) - 2; i >= 0; i-- {
		if primeScores[i+1] > primeScores[i] {
			rightReaches[i] = i + 1
			continue
		}

		j := i + 1
		for j < len(primeScores) && primeScores[j] <= primeScores[i] {
			j = rightReaches[j]
		}
		rightReaches[i] = j
	}
	return rightReaches
}

func generatePrimeScores(nums []int) []int {
	scores := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		scores[i] = calculatePrimeScore(nums[i])
	}
	return scores
}

func calculatePrimeScore(num int) int {
	score := 0
	if num&1 == 0 {
		score++
		for num&1 == 0 {
			num /= 2
		}
	}
	root := int(math.Sqrt(float64(num)))
	for i := 3; i <= root; i += 2 {
		if num%i == 0 {
			score++
			for num%i == 0 {
				num /= i
			}
		}
		if num == 1 {
			break
		}
	}
	if num > 1 {
		score++
	}
	return score
}
