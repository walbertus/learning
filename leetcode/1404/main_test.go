package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Input          string
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Input:          "1",
            ExpectedResult: 0,
        },
        {
            Input:          "10",
            ExpectedResult: 1,
        },
        {
            Input:          "11",
            ExpectedResult: 3,
        },
        {
            Input:          "1101",
            ExpectedResult: 6,
        },
        {
            Input:          "111",
            ExpectedResult: 4,
        },
        {
            Input:          "101",
            ExpectedResult: 5,
        },
        {
            Input:          "11100010100011011",
            ExpectedResult: 26,
        },
    }

    for idx, testcase := range testcases {
        result := numSteps(testcase.Input)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
