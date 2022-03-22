package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase474 struct {
	BinaryStrings   []string
	MaxNumberOfZero int
	MaxNumberOfOne  int
	ExpectedResult  int
}

func Test474(t *testing.T) {
	testcases := []Testcase474{
		{
			BinaryStrings:   []string{"11"},
			MaxNumberOfZero: 1,
			MaxNumberOfOne:  1,
			ExpectedResult:  0,
		},
		{
			BinaryStrings:   []string{"1"},
			MaxNumberOfZero: 1,
			MaxNumberOfOne:  1,
			ExpectedResult:  1,
		},
		{
			BinaryStrings:   []string{"10", "0001", "111001", "1", "0"},
			MaxNumberOfZero: 5,
			MaxNumberOfOne:  3,
			ExpectedResult:  4,
		},
		{
			BinaryStrings:   []string{"10","0","1"},
			MaxNumberOfZero: 1,
			MaxNumberOfOne:  1,
			ExpectedResult:  2,
		},
	}

	for _, testcase := range testcases {
		result := findMaxForm(testcase.BinaryStrings, testcase.MaxNumberOfZero, testcase.MaxNumberOfOne)
		assert.Equal(t, testcase.ExpectedResult, result)
	}
}
