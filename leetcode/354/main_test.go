package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Envelopes      [][]int
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Envelopes:      [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
            ExpectedResult: 3,
        },
        {
            Envelopes:      [][]int{{1, 1}, {1, 1}, {1, 1}},
            ExpectedResult: 1,
        },
        {
            Envelopes:      [][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}},
            ExpectedResult: 4,
        },
        {
            Envelopes:      [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {5, 5}, {6, 7}, {7, 8}},
            ExpectedResult: 7,
        },
        {
            Envelopes:      [][]int{{1, 2}, {2, 3}, {3, 4}, {3, 5}, {4, 5}, {5, 5}, {5, 6}, {6, 7}, {7, 8}},
            ExpectedResult: 7,
        },
    }

    for idx, testcase := range testcases {
        result := maxEnvelopes(testcase.Envelopes)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
