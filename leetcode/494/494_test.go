package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase991 struct {
	NumberSelection []int
	Target          int
	ExpectedResult  int
}

func Test991(t *testing.T) {
	testcases := []Testcase991{
		{
			NumberSelection: []int{100},
			Target:          1,
			ExpectedResult:  0,
		},
		{
			NumberSelection: []int{1},
			Target:          1,
			ExpectedResult:  1,
		},
		{
			NumberSelection: []int{1, 1, 1, 1, 1},
			Target:          3,
			ExpectedResult:  5,
		},
	}

	for idx, testcase := range testcases {
		result := findTargetSumWays(testcase.NumberSelection, testcase.Target)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx)
	}
}
