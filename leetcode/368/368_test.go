package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testcase368 struct {
	nums     []int
	expected []int
}

func Test368(t *testing.T) {
	testcases := []Testcase368{
		{
			nums:     []int{1, 2, 3},
			expected: []int{1, 2},
		},
		{
			nums:     []int{1, 2, 4, 8},
			expected: []int{1, 2, 4, 8},
		},
	}
	for _, tc := range testcases {
		assert.Equalf(t, tc.expected, largestDivisibleSubset(tc.nums), "testcase: %v", tc)
	}
}
