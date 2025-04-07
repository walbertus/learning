package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testcase416 struct {
	nums     []int
	expected bool
}

func Test416(t *testing.T) {
	testcases := []Testcase416{
		{
			nums:     []int{1, 5, 11, 5},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 5},
			expected: false,
		},
	}
	for _, tc := range testcases {
		assert.Equalf(t, tc.expected, canPartition(tc.nums), "testcase: %v", tc)
	}
}
