package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Input          int
    ExpectedResult int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Input:          4,
            ExpectedResult: 8,
        },
        {
            Input:          2,
            ExpectedResult: 2,
        },
        {
            Input:          1,
            ExpectedResult: 1,
        },
        {
            Input:          10,
            ExpectedResult: 700,
        },
        {
            Input:          8,
            ExpectedResult: 132,
        },
        {
            Input:          7,
            ExpectedResult: 41,
        },
    }

    for idx, testcase := range testcases {
        result := countArrangement(testcase.Input)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}

func TestBitmask(t *testing.T) {
    t.Run("isAvailable", func(t *testing.T) {
        testcases := []struct {
            initial  int
            query    int
            expected bool
        }{
            {
                initial:  0,
                query:    1,
                expected: true,
            },
            {
                initial:  0,
                query:    2,
                expected: true,
            },
            {
                initial:  2,
                query:    1,
                expected: false,
            },
            {
                initial:  4,
                query:    2,
                expected: false,
            },
        }
        for _, testcase := range testcases {
            bit := bitmask(testcase.initial)
            assert.Equal(t, testcase.expected, bit.isAvailable(testcase.query))
        }
    })
    t.Run("use", func(t *testing.T) {
        testcases := []struct {
            query    int
            expected bitmask
        }{
            {
                query:    1,
                expected: bitmask(2),
            },
            {
                query:    4,
                expected: bitmask(16),
            },
        }
        for _, testcase := range testcases {
            bit := bitmask(0)
            assert.Equal(t, testcase.expected, bit.use(testcase.query))
        }
    })
}
