package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testcase2818 struct {
	nums     []int
	k        int
	expected int
}

func Test2818(t *testing.T) {
	testcases := []Testcase2818{
		{
			nums:     []int{8, 3, 9, 3, 8},
			k:        2,
			expected: 81,
		},
		{
			nums:     []int{19, 12, 14, 6, 10, 18},
			k:        3,
			expected: 4788,
		},
		{
			nums:     []int{3289, 2832, 14858, 22011},
			k:        6,
			expected: 256720975,
		},
		{
			nums:     []int{1, 7, 11, 1, 5},
			k:        4,
			expected: 9317,
		},
	}

	for idx, testcase := range testcases {
		result := maximumScore(testcase.nums, testcase.k)
		assert.Equalf(t, testcase.expected, result, "testcase: %d", idx)
	}
}

func TestCalculatePrimeScore(t *testing.T) {
	testcases := []struct {
		num      int
		expected int
	}{
		{
			num:      3,
			expected: 1,
		},
		{
			num:      9,
			expected: 1,
		},
		{
			num:      6,
			expected: 2,
		},
		{
			num:      8,
			expected: 1,
		},
		{
			num:      3289,
			expected: 3,
		},
		{
			num:      2832,
			expected: 3,
		},
		{
			num:      14858,
			expected: 4,
		},
		{
			num:      22011,
			expected: 4,
		},
	}

	for idx, testcase := range testcases {
		result := calculatePrimeScore(testcase.num)
		assert.Equalf(t, testcase.expected, result, "testcase: %d for value %d", idx, testcase.num)
	}
}
