package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase struct {
	Input int
	ExpectedResult int
}

func Test(t *testing.T) {
	testcases := []Testcase{
		{
			Input: 1,
			ExpectedResult: 1,
		},
		{
			Input: 2,
			ExpectedResult: 2,
		},
		{
			Input: 5,
			ExpectedResult: 2,
		},
		{
			Input: 13,
			ExpectedResult: 2,
		},
		{
			Input: 12,
			ExpectedResult: 3,
		},
	}

	for idx, testcase := range testcases {
		result := numSquares(testcase.Input)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
	}
}
