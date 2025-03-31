package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testcase2551 struct {
	weights  []int
	k        int
	expected int64
}

func Test2551(t *testing.T) {
	testcases := []Testcase2551{
		{
			weights:  []int{1},
			k:        1,
			expected: 0,
		},
		{
			weights:  []int{1, 2, 3},
			k:        1,
			expected: 0,
		},
		{
			weights:  []int{1, 3, 5, 1},
			k:        2,
			expected: 4,
		},
		{
			weights:  []int{1, 3},
			k:        2,
			expected: 0,
		},
	}

	for idx, testcase := range testcases {
		result := putMarbles(testcase.weights, testcase.k)
		assert.Equalf(t, testcase.expected, result, "testcase: %d, %v", idx, testcase)
	}
}
