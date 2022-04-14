package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Input          int
    ExpectedResult [][]int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Input:          1,
            ExpectedResult: [][]int{{1}},
        },
        {
            Input:          3,
            ExpectedResult: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
        },
    }
    for idx, testcase := range testcases {
        result := generateMatrix(testcase.Input)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
