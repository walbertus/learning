package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Matrix         [][]int
    Query          int
    ExpectedResult []int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Matrix: [][]int{
                {1, 1, 0, 0, 0},
                {1, 1, 1, 1, 0},
                {1, 0, 0, 0, 0},
                {1, 1, 0, 0, 0},
                {1, 1, 1, 1, 1},
            },
            Query:          4,
            ExpectedResult: []int{2, 0, 3, 1},
        },
        {
            Matrix: [][]int{
                {1, 0, 0, 0},
                {1, 1, 1, 1},
                {1, 0, 0, 0},
                {1, 0, 0, 0},
            },
            Query:          4,
            ExpectedResult: []int{0, 2, 3, 1},
        },
    }

    for idx, testcase := range testcases {
        result := kWeakestRows(testcase.Matrix, testcase.Query)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}

func TestNumberOfSoldier(t *testing.T) {
    testcases := []struct {
        List           []int
        ExpectedResult int
    }{
        {
            List:           []int{1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
            ExpectedResult: 3,
        },
        {
            List:           []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
            ExpectedResult: 0,
        },
        {
            List:           []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
            ExpectedResult: 10,
        },
    }
    for idx, testcase := range testcases {
        result := numberOfSoldier(testcase.List)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
