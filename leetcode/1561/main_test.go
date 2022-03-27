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
            Piles:          []int{2, 4, 1, 2, 7, 8},
            ExpectedResult: 9,
        },
        {
            Piles:          []int{2, 4, 5},
            ExpectedResult: 4,
        },
        {
            Piles:          []int{9, 8, 7, 6, 5, 1, 2, 3, 4},
            ExpectedResult: 18,
        },
    }

    for idx, testcase := range testcases {
        result := maxCoins(testcase.Piles)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
