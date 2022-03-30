package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Matrix         [][]int
    Target         int
    ExpectedResult bool
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Matrix:         [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
            Target:         3,
            ExpectedResult: true,
        },
        {
            Matrix:         [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
            Target:         13,
            ExpectedResult: false,
        },
    }

    for idx, testcase := range testcases {
        result := searchMatrix(testcase.Matrix, testcase.Target)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
