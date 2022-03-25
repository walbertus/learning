package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Piles          []int
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Piles:          []int{2},
            ExpectedResult: 2,
        },
        {
            Piles:          []int{2, 3},
            ExpectedResult: 5,
        },
        {
            Piles:          []int{2, 3, 100},
            ExpectedResult: 5,
        },
        {
            Piles:          []int{2, 7, 9, 4, 4},
            ExpectedResult: 10,
        },
        {
            Piles:          []int{1, 2, 3, 4, 5, 100},
            ExpectedResult: 104,
        },
        {
            Piles:          []int{8, 6, 9, 1, 7, 9},
            ExpectedResult: 25,
        },
    }

    for idx, testcase := range testcases {
        result := stoneGameII(testcase.Piles)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
