package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Arr                []int
    MaxPartitionMember int
    ExpectedResult     int
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Arr:                []int{1},
            MaxPartitionMember: 1,
            ExpectedResult:     1,
        },
        {
            Arr:                []int{1, 15, 7, 9, 2, 5, 10},
            MaxPartitionMember: 3,
            ExpectedResult:     84,
        },
        {
            Arr:                []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3},
            MaxPartitionMember: 4,
            ExpectedResult:     83,
        },
    }

    for idx, testcase := range testcases {
        result := maxSumAfterPartitioning(testcase.Arr, testcase.MaxPartitionMember)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
