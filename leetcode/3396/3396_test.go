package leetcode

import "testing"

type Testcase3396 struct {
	Nums           []int
	ExpectedResult int
}

func Test3396(t *testing.T) {
	testcases := []Testcase3396{
		{
			Nums:           []int{1, 2, 3, 4, 2, 3, 3, 5, 7},
			ExpectedResult: 2,
		},
		{
			Nums:           []int{4, 5, 6, 4, 4},
			ExpectedResult: 2,
		},
	}
	for idx, testcase := range testcases {
		result := minimumOperations(testcase.Nums)
		if result != testcase.ExpectedResult {
			t.Errorf("Testcase %d failed: expected %d, got %d", idx, testcase.ExpectedResult, result)
		}
	}
}
