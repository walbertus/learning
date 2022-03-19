package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase3 struct {
	Input    string
	Expected int
}

func Test3(t *testing.T) {
	testcases := []Testcase3{
		{Input: "", Expected: 0},
		{Input: "a", Expected: 1},
		{Input: "dvdf", Expected: 3},
		{Input: "asjrgapa", Expected: 6},
		{Input: "abcdef", Expected: 6},
		{Input: "abcabcbb", Expected: 3},
		{Input: "bbbbb", Expected: 1},
		{Input: "pwwkew", Expected: 3},
	}

	for _, testcase := range testcases {
		result := lengthOfLongestSubstring(testcase.Input)
		assert.Equal(t, testcase.Expected, result)
	}
}
