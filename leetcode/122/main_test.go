package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Prices         []int
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Prices:         []int{7, 1, 5, 3, 6, 4},
            ExpectedResult: 7,
        },
        {
            Prices:         []int{1, 2, 3, 4, 5},
            ExpectedResult: 4,
        },
        {
            Prices:         []int{7, 6, 4, 3, 1},
            ExpectedResult: 0,
        },
        {
            Prices:         []int{1, 100, 1, 100},
            ExpectedResult: 198,
        },
    }

    for idx, testcase := range testcases {
        result := maxProfit(testcase.Prices)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
