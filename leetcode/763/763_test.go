package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testcase763 struct {
	Input          string
	ExpectedResult []int
}

func Test763(t *testing.T) {
	testcases := []Testcase763{
		{
			Input:          "a",
			ExpectedResult: []int{1},
		},
		{
			Input:          "ab",
			ExpectedResult: []int{1, 1},
		},
		{
			Input:          "abbacd",
			ExpectedResult: []int{4, 1, 1},
		},
		{
			Input:          "ababcbacadefegdehijhklij",
			ExpectedResult: []int{9, 7, 8},
		},
		{
			Input:          "eccbbbbdec",
			ExpectedResult: []int{10},
		},
	}

	for idx, testcase := range testcases {
		result := partitionLabels(testcase.Input)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx)
	}
}
