package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase991 struct {
	StartValue     int
	Target         int
	ExpectedResult int
}

func Test991(t *testing.T) {
	testcases := []Testcase991{
		{
			StartValue:     1,
			Target:         1,
			ExpectedResult: 0,
		},
		{
			StartValue:     3,
			Target:         1,
			ExpectedResult: 2,
		},
		{
			StartValue:     2,
			Target:         3,
			ExpectedResult: 2,
		},
		{
			StartValue:     5,
			Target:         8,
			ExpectedResult: 2,
		},
		{
			StartValue:     3,
			Target:         10,
			ExpectedResult: 3,
		},
		{
			StartValue:     1,
			Target:         9,
			ExpectedResult: 7,
		},
	}

	for idx, testcase := range testcases {
		result := brokenCalc(testcase.StartValue, testcase.Target)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx)
	}
}
