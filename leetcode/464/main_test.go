package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    MaxChoosableInteger int
    DesiredTotal        int
    ExpectedResult      bool
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        0,
            ExpectedResult:      true,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        10,
            ExpectedResult:      true,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        11,
            ExpectedResult:      false,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        12,
            ExpectedResult:      true,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        13,
            ExpectedResult:      true,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        14,
            ExpectedResult:      true,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        15,
            ExpectedResult:      true,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        20,
            ExpectedResult:      true,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        21,
            ExpectedResult:      true,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        22,
            ExpectedResult:      false,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        33,
            ExpectedResult:      false,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        37,
            ExpectedResult:      false,
        },
        {
            MaxChoosableInteger: 10,
            DesiredTotal:        40,
            ExpectedResult:      false,
        },
        {
            MaxChoosableInteger: 18,
            DesiredTotal:        79,
            ExpectedResult:      true,
        },
    }

    for idx, testcase := range testcases {
        result := canIWin(testcase.MaxChoosableInteger, testcase.DesiredTotal)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}

func TestChoosableInteger(t *testing.T) {
    t.Run("add", func(t *testing.T) {
        testcases := []struct {
            maxChoosableInteger int
            bitmask bitmask
        }{
            {
                maxChoosableInteger: 3,
                bitmask: 14,
            },
            {
                maxChoosableInteger: 6,
                bitmask: 126,
            },
        }
        for _, testcase := range testcases {
            target := bitmask(0)
            for i := 1; i <= testcase.maxChoosableInteger; i++ {
                target = target.add(i)
            }
            assert.Equal(t, int(testcase.bitmask), int(target))
        }
    })
    t.Run("Test getInts", func(t *testing.T) {
        testcases := []struct {
            bitmask bitmask
            ints    []int
        }{
            {
                bitmask: 10,
                ints:    []int{1, 3},
            },
            {
                bitmask: 20,
                ints:    []int{2, 4},
            },
        }
        for _, testcase := range testcases {
            target := testcase.bitmask
            ints := target.getInts()
            assert.Equal(t, testcase.ints, ints)
        }
    })
    t.Run("Test remove", func(t *testing.T) {
        testcases := []struct {
            bitmask bitmask
            remove int
            expectedResult bitmask
        }{
            {
                bitmask: 14,
                remove: 2,
                expectedResult: 10,
            },
            {
                bitmask: 78,
                remove: 6,
                expectedResult: 14,
            },
        }
        for _, testcase := range testcases {
            assert.Equal(t, testcase.expectedResult, testcase.bitmask.remove(testcase.remove))
        }
    })
}
