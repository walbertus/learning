package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Nums           []int
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Nums:           []int{3, 4, 2},
            ExpectedResult: 6,
        },
        {
            Nums:           []int{2, 2, 3, 3, 3, 4},
            ExpectedResult: 9,
        },
    }

    for idx, testcase := range testcases {
        result := deleteAndEarn(testcase.Nums)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
