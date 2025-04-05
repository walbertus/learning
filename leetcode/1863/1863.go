package leetcode

const BIT_LIMIT = 4

// https://leetcode.com/problems/sum-of-all-subset-xor-totals
// TODO: It seems this solution is overcomplicated
func subsetXORSum(nums []int) int {
	bitCounting := countBitsOfNums(nums, BIT_LIMIT)
	total := 0
	for i := 0; i <= BIT_LIMIT; i++ {
		if bitCounting[i] == 0 {
			continue
		}

		numberOfOne := bitCounting[i]
		numberOfZero := len(nums) - numberOfOne

		for j := 1; j <= len(nums); j += 2 {
			for k := 0; k <= len(nums)-j; k++ {
				total += calculateXOREqualOne(numberOfOne, numberOfZero, j, k) * (1 << i)
			}
		}

	}
	return total
}

func calculateXOREqualOne(numberOfOne int, numberOfZero int, sizeOne int, sizeZero int) int {
	result := 0
	result += calculateCombinatoric(numberOfOne, sizeOne) * calculateCombinatoric(numberOfZero, sizeZero)
	return result
}

func calculateCombinatoric(n int, k int) int {
	if n < k {
		return 0
	}
	if n == k {
		return 1
	}
	if k == 0 {
		return 1
	}
	if k > n/2 {
		k = n - k
	}

	result := 1
	for i := 0; i < k; i++ {
		result *= (n - i)
		result /= (i + 1)
	}
	return result
}

func countBitsOfNums(nums []int, bitLimit int) []int {
	bitCounting := make([]int, bitLimit+1)
	for _, num := range nums {
		for i := 0; i <= bitLimit; i++ {
			if (1<<i)&num != 0 {
				bitCounting[i]++
			}
		}
	}
	return bitCounting
}
