package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase6 struct {
	InputString    string
	InputRow       int
	ExpectedResult string
}

func Test6(t *testing.T) {
	testcases := []Testcase6{
		{
			InputString: "ABC",
			InputRow: 1,
			ExpectedResult: "ABC",
		},
		{
			InputString:    "PAYPALISHIRING",
			InputRow:       3,
			ExpectedResult: "PAHNAPLSIIGYIR",
		},
		{
			InputString:    "PAYPALISHIRING",
			InputRow:       4,
			ExpectedResult: "PINALSIGYAHRPI",
		},
		{
			InputString:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			InputRow:       3,
			ExpectedResult: "AEIMQUYBDFHJLNPRTVXZCGKOSW",
		},
		{
			InputString:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			InputRow:       4,
			ExpectedResult: "AGMSYBFHLNRTXZCEIKOQUWDJPV",
		},
		{
			InputString:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			InputRow:       5,
			ExpectedResult: "AIQYBHJPRXZCGKOSWDFLNTVEMU",
		},
	}

	for _, testcase := range testcases {
		result := convert(testcase.InputString, testcase.InputRow)
		assert.Equal(t, testcase.ExpectedResult, result)
	}
}
