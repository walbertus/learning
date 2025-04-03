package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testcase1406 struct {
	stoneValue []int
	expected   string
}

func Test1406(t *testing.T) {
	testcase := []Testcase1406{
		{
			stoneValue: []int{1},
			expected:   "Alice",
		},
		{
			stoneValue: []int{1, 2, 3},
			expected:   "Alice",
		},
		{
			stoneValue: []int{-1, -2, -7},
			expected:   "Alice",
		},
		{
			stoneValue: []int{1, 2, 3, 7},
			expected:   "Bob",
		},
		{
			stoneValue: []int{1, 2, 3, -9},
			expected:   "Alice",
		},
		{
			stoneValue: []int{1, 2, 3, 6},
			expected:   "Tie",
		},
		{
			stoneValue: []int{-1, -2, -3},
			expected:   "Tie",
		},
	}
	for idx, testcase := range testcase {
		result := stoneGameIII(testcase.stoneValue)
		assert.Equalf(t, testcase.expected, result, "testcase: %d", idx+1)
	}
}
