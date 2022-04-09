package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Nums           []int
    K              int
    ExpectedResult []int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Nums:           []int{1, 2},
            K:              2,
            ExpectedResult: []int{1, 2},
        },
        {
            Nums:           []int{1},
            K:              1,
            ExpectedResult: []int{1},
        },
        {
            Nums:           []int{-1, -1},
            K:              1,
            ExpectedResult: []int{-1},
        },
        {
            Nums:           []int{1, 1, 1, 2, 2, 3},
            K:              2,
            ExpectedResult: []int{1, 2},
        },
    }

    for idx, testcase := range testcases {
        result := topKFrequent(testcase.Nums, testcase.K)
        assert.ElementsMatchf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
