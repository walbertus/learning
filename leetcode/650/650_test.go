package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase200 struct {
	Target         int
	ExpectedResult int
}

func Test200(t *testing.T) {
	testcases := []Testcase200{
		{
			Target:         3,
			ExpectedResult: 3,
		},
		{
			Target:         10,
			ExpectedResult: 7,
		},
		{
			Target:         1000,
			ExpectedResult: 21,
		},
		{
			Target:         1,
			ExpectedResult: 0,
		},
	}

	for idx, testcase := range testcases {
		result := minSteps(testcase.Target)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx)
	}
}
