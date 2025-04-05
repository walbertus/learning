package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testcase2863 struct {
	nums     []int
	expected int
}

func Test1863(t *testing.T) {
	testcases := []Testcase2863{
		{
			nums:     []int{1, 3},
			expected: 6,
		},
		{
			nums:     []int{1, 1, 1},
			expected: 4,
		},
		{
			nums:     []int{5, 1, 6},
			expected: 28,
		},
		{
			nums:     []int{3, 4, 5, 6, 7, 8},
			expected: 480,
		},
		{
			nums:     []int{5, 17, 17, 12, 17},
			expected: 464,
		},
	}

	for idx, testcase := range testcases {
		result := subsetXORSum(testcase.nums)
		assert.Equalf(t, testcase.expected, result, "testcase: %d", idx)
	}
}
func TestCountBitsOfNums(t *testing.T) {
	testcases := []struct {
		nums     []int
		bitLimit int
		expected []int
	}{
		{
			nums:     []int{1, 3},
			bitLimit: 3,
			expected: []int{2, 1, 0, 0},
		},
		{
			nums:     []int{5, 1, 6},
			bitLimit: 3,
			expected: []int{2, 1, 2, 0},
		},
		{
			nums:     []int{3, 4, 5, 6, 7, 8},
			bitLimit: 3,
			expected: []int{3, 3, 4, 1},
		},
		{
			nums:     []int{1, 1, 1},
			bitLimit: 3,
			expected: []int{3, 0, 0, 0},
		},
	}

	for idx, testcase := range testcases {
		result := countBitsOfNums(testcase.nums, testcase.bitLimit)
		assert.Equalf(t, testcase.expected, result, "testcase: %d", idx)
	}
}

func TestCalculateXOREqualOne(t *testing.T) {
	testcases := []struct {
		numberOfOne  int
		numberOfZero int
		sizeOne      int
		sizeZero     int
		expected     int
	}{
		{
			numberOfOne:  3,
			numberOfZero: 3,
			sizeOne:      3,
			sizeZero:     0,
			expected:     1,
		},
		{
			numberOfOne:  3,
			numberOfZero: 3,
			sizeOne:      1,
			sizeZero:     0,
			expected:     3,
		},
		{
			numberOfOne:  3,
			numberOfZero: 3,
			sizeOne:      1,
			sizeZero:     1,
			expected:     9,
		},
	}

	for idx, testcase := range testcases {
		result := calculateXOREqualOne(testcase.numberOfOne, testcase.numberOfZero, testcase.sizeOne, testcase.sizeZero)
		assert.Equalf(t, testcase.expected, result, "testcase: %d", idx)
	}
}
