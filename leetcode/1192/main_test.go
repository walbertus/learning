package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    NumberOfNode   int
    Edges          [][]int
    ExpectedResult [][]int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            NumberOfNode:   4,
            Edges:          [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}},
            ExpectedResult: [][]int{{1, 3}},
        },
    }

    for idx, testcase := range testcases {
        result := criticalConnections(testcase.NumberOfNode, testcase.Edges)
        assert.ElementsMatchf(t, result, testcase.ExpectedResult, "testcase: %d", idx+1)
    }
}
