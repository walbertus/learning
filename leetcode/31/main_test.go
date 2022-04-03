package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Input          []int
    ExpectedResult []int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Input:          []int{1},
            ExpectedResult: []int{1},
        },
        {
            Input:          []int{3, 2, 1},
            ExpectedResult: []int{1, 2, 3},
        },
        {
            Input:          []int{1, 2, 3},
            ExpectedResult: []int{1, 3, 2},
        },
        {
            Input:          []int{1, 1, 5},
            ExpectedResult: []int{1, 5, 1},
        },
        {
            Input:          []int{4, 2, 3, 1},
            ExpectedResult: []int{4, 3, 1, 2},
        },
        {
            Input:          []int{1, 4, 3, 2},
            ExpectedResult: []int{2, 1, 3, 4},
        },
        {
            Input:          []int{2, 3, 1, 3, 3},
            ExpectedResult: []int{2, 3, 3, 1, 3},
        },
    }

    for idx, testcase := range testcases {
        nextPermutation(testcase.Input)
        assert.Equalf(t, testcase.ExpectedResult, testcase.Input, "testcase: %d", idx+1)
    }
}
