package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase3394 struct {
	n          int
	rectangles [][]int
	expected   bool
}

func Test3394(t *testing.T) {
	testcases := []TestCase3394{
		{
			n:          5,
			rectangles: [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}},
			expected:   true,
		},
		{
			n:          4,
			rectangles: [][]int{{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}},
			expected:   true,
		},
		{
			n:          4,
			rectangles: [][]int{{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}},
			expected:   false,
		},
		{
			n:          5,
			rectangles: [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}},
			expected:   true,
		},
		{
			n:          3,
			rectangles: [][]int{{0, 0, 1, 2}, {1, 0, 2, 2}, {2, 0, 3, 2}, {0, 2, 1, 3}, {1, 2, 3, 3}},
			expected:   false,
		},
		{
			n:          6,
			rectangles: [][]int{{4, 0, 5, 2}, {0, 5, 4, 6}, {4, 5, 6, 6}},
			expected:   false,
		},
	}

	for idx, testcase := range testcases {
		result := checkValidCuts(testcase.n, testcase.rectangles)
		assert.Equalf(t, testcase.expected, result, "testcase: %d", idx)
	}
}
func TestLineSweep(t *testing.T) {
	testcases := []struct {
		counter  map[int]int
		expected int
	}{
		{
			counter:  map[int]int{0: 0, 1: 0, 2: 1, 3: -1},
			expected: 2,
		},
		{
			counter:  map[int]int{1: 1, 5: 1, 4: -1, 6: -1},
			expected: 1,
		},
		{
			counter:  map[int]int{1: 1, 2: -1, 6: 0},
			expected: 1,
		},
	}

	for idx, testcase := range testcases {
		result := lineSweep(testcase.counter)
		assert.Equalf(t, testcase.expected, result, "testcase: %d", idx)
	}
}
