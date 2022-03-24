package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase174 struct {
	DungeonData    [][]int
	ExpectedResult int
}

func Test174(t *testing.T) {
	testcases := []Testcase174{
		{
			DungeonData: [][]int{
				{
					-2, -3, 3,
				}, {
					-5, -10, 1,
				}, {
					10, 30, -5,
				},
			},
			ExpectedResult: 7,
		},
		{
			DungeonData: [][]int{
				{
					-2, -3,
				}, {
					-5, -10,
				},
			},
			ExpectedResult: 16,
		},
		{
			DungeonData: [][]int{
				{
					-2, 100,
				}, {
					-5, -10,
				},
			},
			ExpectedResult: 3,
		},
		{
			DungeonData: [][]int{
				{
					0,
				},
			},
			ExpectedResult: 1,
		},
		{
			DungeonData: [][]int{
				{
					-5,
				},
			},
			ExpectedResult: 6,
		},
		{
			DungeonData: [][]int{
				{
					6,
				},
			},
			ExpectedResult: 1,
		},
	}

	for idx, testcase := range testcases {
		result := calculateMinimumHP(testcase.DungeonData)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx)
	}
}
