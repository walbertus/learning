package leetcode

import "testing"

type Testcase2874 struct {
	nums     []int
	expected int64
}

func Test2874(t *testing.T) {
	testcases := []Testcase2874{
		{
			nums:     []int{12, 6, 1, 2, 7},
			expected: 77,
		},
		{
			nums:     []int{1, 10, 3, 4, 19},
			expected: 133,
		},
		{
			nums:     []int{1, 2, 3},
			expected: 0,
		},
		{
			nums:     []int{10, 1, 10},
			expected: 90,
		},
		{
			nums:     []int{6, 11, 12, 12, 7, 9, 2, 11, 12, 4, 19, 14, 16, 8, 16},
			expected: 190,
		},
	}

	for idx, testcase := range testcases {
		result := maximumTripletValue(testcase.nums)
		if result != testcase.expected {
			t.Errorf("testcase: %d, expected: %d, got: %d", idx, testcase.expected, result)
		}
	}
}
