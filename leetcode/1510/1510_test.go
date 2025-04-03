package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testcase1510 struct {
	n        int
	expected bool
}

func Test1510(t *testing.T) {
	testcases := []Testcase1510{
		// {
		// 	n:        1,
		// 	expected: true,
		// },
		// {
		// 	n:        2,
		// 	expected: false,
		// },
		// {
		// 	n:        3,
		// 	expected: true,
		// },
		// {
		// 	n:        4,
		// 	expected: true,
		// },
		// {
		// 	n:        5,
		// 	expected: false,
		// },
		// {
		// 	n:        17,
		// 	expected: false,
		// },
		// {
		// 	n:        154,
		// 	expected: true,
		// },
		{
			n:        4676,
			expected: true,
		},
	}

	for idx, testcase := range testcases {
		result := winnerSquareGame(testcase.n)
		if result != testcase.expected {
			assert.Equalf(t, testcase.expected, result, "testcase: %d", idx)
		}
	}
}
