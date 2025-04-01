package leetcode

import "testing"

type Testcase2140 struct {
	questions [][]int
	expected  int64
}

func Test2140(t *testing.T) {
	testcase := []Testcase2140{
		{
			questions: [][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}},
			expected:  5,
		},
		{
			questions: [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}},
			expected:  7,
		},
		{
			questions: [][]int{{12, 46}, {78, 19}, {63, 15}, {79, 62}, {13, 10}},
			expected:  79,
		},
	}

	for idx, testcase := range testcase {
		result := mostPoints(testcase.questions)
		if result != testcase.expected {
			t.Errorf("testcase: %d, expected: %d, got: %d", idx, testcase.expected, result)
		}
	}
}
