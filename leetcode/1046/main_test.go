package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Input          []int
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Input:          []int{2, 7, 4, 1, 8, 1},
            ExpectedResult: 1,
        },
        {
            Input:          []int{1},
            ExpectedResult: 1,
        },
    }

    for idx, testcase := range testcases {
        result := lastStoneWeight(testcase.Input)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
