package leetcode

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Testcase struct {
    Equations      [][]string
    Values         []float64
    Queries        [][]string
    ExpectedResult []float64
}

func Test(t *testing.T) {
    testcases := []Testcase{
        {
            Equations:      [][]string{{"a", "b"}, {"b", "c"}, {"d", "e"}},
            Values:         []float64{2.0, 3.0, 1.0},
            Queries:        [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}, {"a", "d"}},
            ExpectedResult: []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000, -1},
        },
        {
            Equations:      [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
            Values:         []float64{1.5, 2.5, 5.0},
            Queries:        [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
            ExpectedResult: []float64{3.75000, 0.40000, 5.00000, 0.20000},
        },
    }

    for idx, testcase := range testcases {
        result := calcEquation(testcase.Equations, testcase.Values, testcase.Queries)
        assert.Equalf(t, testcase.ExpectedResult, result, "testcase: %d", idx+1)
    }
}
