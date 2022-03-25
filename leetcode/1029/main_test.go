package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase struct {
	Input          [][]int
	ExpectedResult int
}

func Test(t *testing.T) {
	testcases := []Testcase{
		{
			Input: [][]int{
				{259, 770},
				{448, 54},
				{926, 667},
				{184, 139},
				{840, 118},
				{577, 469},
			},
			ExpectedResult: 1859,
		},
	}

	for idx, testcase := range testcases {
		result := twoCitySchedCost(testcase.Input)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
	}
}
