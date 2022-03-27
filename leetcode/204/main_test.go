package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Max            int
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Max:            10,
            ExpectedResult: 4,
        },
        {
            Max:            0,
            ExpectedResult: 0,
        },
        {
            Max:            1,
            ExpectedResult: 0,
        },
        {
            Max:            30,
            ExpectedResult: 10,
        },
    }

    for idx, testcase := range testcases {
        result := countPrimes(testcase.Max)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
