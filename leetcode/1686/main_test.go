package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Alice          []int
    Bob            []int
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Alice:          []int{1, 3},
            Bob:            []int{2, 1},
            ExpectedResult: 1,
        },
        {
            Alice:          []int{1, 2},
            Bob:            []int{3, 1},
            ExpectedResult: 0,
        },
        {
            Alice:          []int{2, 4, 3},
            Bob:            []int{1, 6, 7},
            ExpectedResult: -1,
        },
        {
            Alice:          []int{1, 48, 7, 86, 91},
            Bob:            []int{83, 99, 51, 41, 42},
            ExpectedResult: 1,
        },
    }

    for idx, testcase := range testcases {
        result := stoneGameVI(testcase.Alice, testcase.Bob)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
