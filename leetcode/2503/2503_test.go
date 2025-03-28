package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testcase2503 struct {
	grid     [][]int
	queries  []int
	expected []int
}

func Test2503(t *testing.T) {
	testcases := []Testcase2503{
		{
			grid:     [][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}},
			queries:  []int{5, 6, 2},
			expected: []int{5, 8, 1},
		},
		{
			grid:     [][]int{{5, 2, 1}, {1, 1, 2}},
			queries:  []int{3},
			expected: []int{0},
		},
		{
			grid:     [][]int{{1, 1, 5}, {1, 5, 1}, {5, 1, 1}},
			queries:  []int{2, 5, 6},
			expected: []int{3, 3, 9},
		},
		{
			grid: [][]int{
				{123491, 95183, 131119, 576084, 779700, 886039, 564610},
				{835246, 594630, 752204, 976312, 431928, 916878, 37773},
				{602559, 675, 8018, 72760, 560850, 132858, 416126},
				{787316, 77587, 784798, 797907, 769783, 143785, 378185},
				{362862, 754648, 212843, 813454, 552332, 10700, 266493},
			},
			queries:  []int{581002, 174698},
			expected: []int{4, 3},
		},
	}

	for idx, testcase := range testcases {
		result := maxPoints(testcase.grid, testcase.queries)
		assert.Equalf(t, testcase.expected, result, "testcase: %d", idx)
	}
}
