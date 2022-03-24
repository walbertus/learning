package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase struct {
	PeopleWeight   []int
	BoatLimit      int
	ExpectedResult int
}

func Test(t *testing.T) {
	testcases := []Testcase{
		{
			PeopleWeight:   []int{1, 2},
			BoatLimit:      3,
			ExpectedResult: 1,
		},
		{
			PeopleWeight:   []int{5, 5, 2, 1},
			BoatLimit:      3,
			ExpectedResult: 3,
		},
		{
			PeopleWeight:   []int{3, 5, 3, 4},
			BoatLimit:      5,
			ExpectedResult: 4,
		},
	}

	for idx, testcase := range testcases {
		result := numRescueBoats(testcase.PeopleWeight, testcase.BoatLimit)
		assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
	}
}
