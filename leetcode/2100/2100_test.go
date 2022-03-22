package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase2100 struct {
	Securities     []int
	SafeTime       int
	ExpectedResult []int
}

func Test2100(t *testing.T) {
	testcases := []Testcase2100{
		{
			Securities:     []int{5, 3, 3, 3, 5, 6, 2},
			SafeTime:       2,
			ExpectedResult: []int{2, 3},
		},
		{
			Securities:     []int{1, 1, 1, 1, 1},
			SafeTime:       0,
			ExpectedResult: []int{0, 1, 2, 3, 4},
		},
		{
			Securities:     []int{1, 2, 3, 4, 5, 6},
			SafeTime:       2,
			ExpectedResult: []int{},
		},
	}

	for idx, testcase := range testcases {
		result := goodDaysToRobBank(testcase.Securities, testcase.SafeTime)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx)
	}
}
