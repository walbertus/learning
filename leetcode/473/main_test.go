package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Input          []int
    ExpectedResult bool
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Input:          []int{1, 1, 2, 2, 2},
            ExpectedResult: true,
        },
        {
            Input:          []int{3, 3, 3, 3, 4},
            ExpectedResult: false,
        },
        {
            Input:          []int{1, 1, 2, 2, 2, 4, 4, 4, 4},
            ExpectedResult: true,
        },
        {
            Input:          []int{5, 5, 5, 5, 16, 4, 4, 4, 4, 4, 3, 3, 3, 3, 4},
            ExpectedResult: false,
        },
    }

    for idx, testcase := range testcases {
        result := makesquare(testcase.Input)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
