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
            Input:          [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
            ExpectedResult: 2,
        },
        {
            Input:          [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}},
            ExpectedResult: 1,
        },
        {
            Input:          [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}},
            ExpectedResult: 0,
        },
    }

    for idx, testcase := range testcases {
        result := minimumEffortPath(testcase.Input)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
