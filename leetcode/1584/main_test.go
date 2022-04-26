package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Input          [][]int
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Input:          [][]int{{0, 0}},
            ExpectedResult: 0,
        },
        {
            Input:          [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
            ExpectedResult: 20,
        },
        {
            Input:          [][]int{{-14, -14}, {18, -10}, {18, 18}, {10, -2}},
            ExpectedResult: 80,
        },
    }

    for idx, testcase := range testcases {
        result := minCostConnectPoints(testcase.Input)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}

func TestGenerateEdgeList(t *testing.T) {
    testcases := []struct {
        points   [][]int
        expected []Edge
    }{
        {
            points: [][]int{{0, 1}, {10, 10}, {3, 3}},
            expected: []Edge{
                {
                    source:      0,
                    destination: 1,
                    cost:        19,
                }, {
                    source:      0,
                    destination: 2,
                    cost:        5,
                }, {
                    source:      1,
                    destination: 2,
                    cost:        14,
                },
            },
        }, {
            points: [][]int{{-14, -14}, {-18, 5}, {18, -10}, {18, 18}, {10, -2}},
            expected: []Edge{
                {
                    source:      0,
                    destination: 1,
                    cost:        23,
                }, {
                    source:      0,
                    destination: 2,
                    cost:        36,
                }, {
                    source:      0,
                    destination: 3,
                    cost:        64,
                }, {
                    source:      0,
                    destination: 4,
                    cost:        36,
                }, {
                    source:      1,
                    destination: 2,
                    cost:        51,
                }, {
                    source:      1,
                    destination: 3,
                    cost:        49,
                }, {
                    source:      1,
                    destination: 4,
                    cost:        35,
                }, {
                    source:      2,
                    destination: 3,
                    cost:        28,
                }, {
                    source:      2,
                    destination: 4,
                    cost:        16,
                }, {
                    source:      3,
                    destination: 4,
                    cost:        28,
                },
            },
        },
    }
    for idx, testcase := range testcases {
        result := generateSorterEdgeList(testcase.points)
        assert.ElementsMatchf(t, testcase.expected, result, "testcase: %d", idx+1)
    }
}
