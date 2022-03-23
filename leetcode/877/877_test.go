package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase877 struct {
	Piles          []int
	ExpectedResult bool
}

func Test877(t *testing.T) {
	testcases := []Testcase877{
		{
			Piles:          []int{5, 3},
			ExpectedResult: true,
		},
		{
			Piles:          []int{1, 10, 100, 101, 10, 1},
			ExpectedResult: true,
		},
		{
			Piles:          []int{5, 3, 4, 5},
			ExpectedResult: true,
		},
		{
			Piles:          []int{3, 7, 2, 3},
			ExpectedResult: true,
		},
	}

	for idx, testcase := range testcases {
		result := stoneGame(testcase.Piles)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
	}
}
