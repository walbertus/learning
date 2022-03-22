package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase1663 struct {
	LengthOfString   int
	SumOfStringValue int
	ExpectedResult   string
}

func Test1663(t *testing.T) {
	testcases := []Testcase1663{
		{
			LengthOfString:   1,
			SumOfStringValue: 1,
			ExpectedResult:   "a",
		},
		{
			LengthOfString:   2,
			SumOfStringValue: 2,
			ExpectedResult:   "aa",
		},
		{
			LengthOfString:   2,
			SumOfStringValue: 52,
			ExpectedResult:   "zz",
		},
		{
			LengthOfString:   3,
			SumOfStringValue: 27,
			ExpectedResult:   "aay",
		},
		{
			LengthOfString:   5,
			SumOfStringValue: 73,
			ExpectedResult:   "aaszz",
		},
	}

	for _, testcase := range testcases {
		result := getSmallestString(testcase.LengthOfString, testcase.SumOfStringValue)
		assert.Equal(t, testcase.ExpectedResult, result)
	}
}
