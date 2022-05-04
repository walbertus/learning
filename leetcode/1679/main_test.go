package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Nums           []int
    K              int
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Nums:           []int{1, 2, 3, 4},
            K:              5,
            ExpectedResult: 2,
        },
        {
            Nums:           []int{1, 4, 3, 2},
            K:              5,
            ExpectedResult: 2,
        },
        {
            Nums:           []int{3, 1, 3, 4, 3},
            K:              6,
            ExpectedResult: 1,
        },
        {
            Nums:           []int{2, 2, 2, 3, 1, 1, 4, 1},
            K:              4,
            ExpectedResult: 2,
        },
    }

    for idx, testcase := range testcases {
        result := maxOperations(testcase.Nums, testcase.K)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
