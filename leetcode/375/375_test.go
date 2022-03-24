package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase struct {
	NumberLimit    int
	ExpectedResult int
}

func Test(t *testing.T) {
	testcases := []Testcase{
		{
			NumberLimit:    1,
			ExpectedResult: 0,
		},
		{
			NumberLimit:    2,
			ExpectedResult: 1,
		},
		{
			NumberLimit:    3,
			ExpectedResult: 2,
		},
		{
			NumberLimit:    4,
			ExpectedResult: 4,
		},
		{
			NumberLimit:    10,
			ExpectedResult: 16,
		},
	}

	for idx, testcase := range testcases {
		result := getMoneyAmount(testcase.NumberLimit)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
	}
}
