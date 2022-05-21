package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Coins          []int
    Amount         int
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Coins:          []int{1},
            Amount:         0,
            ExpectedResult: 0,
        },
        {
            Coins:          []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422},
            Amount:         9864,
            ExpectedResult: 24,
        },
        {
            Coins:          []int{1, 2, 5},
            Amount:         11,
            ExpectedResult: 3,
        },
    }

    for idx, testcase := range testcases {
        result := coinChange(testcase.Coins, testcase.Amount)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
