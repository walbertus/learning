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
            Input:          []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
            ExpectedResult: 49,
        },
        {
            Input:          []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
            ExpectedResult: 49,
        },
        {
            Input:          []int{1, 2, 1},
            ExpectedResult: 2,
        },
        {
            Input:          []int{0, 3, 4, 3, 0},
            ExpectedResult: 6,
        },
        {
            Input:          []int{2, 1, 1, 1, 1, 1, 5, 4},
            ExpectedResult: 14,
        },
    }

    for idx, testcase := range testcases {
        result := maxArea(testcase.Input)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
