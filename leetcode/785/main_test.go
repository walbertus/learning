package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Input          [][]int
    ExpectedResult bool
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Input:          [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}},
            ExpectedResult: false,
        },
        {
            Input:          [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}},
            ExpectedResult: true,
        },
    }

    for idx, testcase := range testcases {
        result := isBipartite(testcase.Input)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
