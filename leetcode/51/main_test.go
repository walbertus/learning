package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Input          int
    ExpectedResult [][]string
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Input:          1,
            ExpectedResult: [][]string{{"Q"}},
        },
        {
            Input:          2,
            ExpectedResult: [][]string{},
        },
        {
            Input:          3,
            ExpectedResult: [][]string{},
        },
        {
            Input:          4,
            ExpectedResult: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
        },
    }

    for idx, testcase := range testcases {
        result := solveNQueens(testcase.Input)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}

func TestConvertBoard(t *testing.T) {
    testcases := []struct {
        board          [][]int
        size           int
        expectedResult []string
    }{
        {
            board:          [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
            size:           3,
            expectedResult: []string{"Q..", ".Q.", "..Q"},
        },
    }

    for _, testcase := range testcases {
        result := convertBoard(testcase.board, testcase.size)
        assert.Equal(t, testcase.expectedResult, result)
    }
}
