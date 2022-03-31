package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Nums           []int
    ExpectedResult bool
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Nums:           []int{1, 5, 2},
            ExpectedResult: false,
        },
        {
            Nums:           []int{1, 5, 233, 7},
            ExpectedResult: true,
        },
        {
            Nums:           []int{1, 2, 99},
            ExpectedResult: true,
        },
        {
            Nums:           []int{1, 3, 1},
            ExpectedResult: false,
        },
    }

    for idx, testcase := range testcases {
        result := PredictTheWinner(testcase.Nums)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
