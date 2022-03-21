package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase200 struct {
	Grid           [][]byte
	ExpectedResult int
}

func Test200(t *testing.T) {
	testcases := []Testcase200{
		{
			Grid: [][]byte{
				{'0'},
			},
			ExpectedResult: 0,
		},
		{
			Grid: [][]byte{
				{'1'},
			},
			ExpectedResult: 1,
		},
		{
			Grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			ExpectedResult: 1,
		},
		{
			Grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			ExpectedResult: 3,
		},
	}

	for idx, testcase := range testcases {
		result := numIslands(testcase.Grid)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx)
	}
}
