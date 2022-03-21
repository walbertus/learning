package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testcase1357 struct {
	Constructor ConstructorTestObject
	Operations  []OperationTestObject
}

type ConstructorTestObject struct {
	DiscountCheckpoint int
	Discount           int
	Products           []int
	Prices             []int
}

type OperationTestObject struct {
	Products       []int
	Amount         []int
	ExpectedResult float64
}

func Test1357(t *testing.T) {
	testcases := []Testcase1357{
		{
			Constructor: ConstructorTestObject{
				DiscountCheckpoint: 3,
				Discount:           50,
				Products:           []int{1, 2, 3, 4, 5, 6, 7, 8},
				Prices:             []int{100, 200, 300, 400, 300, 200, 100, 3},
			},
			Operations: []OperationTestObject{
				{
					Products:       []int{1, 2},
					Amount:         []int{1, 2},
					ExpectedResult: 500.0,
				},
				{
					Products:       []int{3, 7},
					Amount:         []int{10, 10},
					ExpectedResult: 4000.0,
				},
				{
					Products:       []int{1, 2, 3, 4, 5, 6, 7},
					Amount:         []int{1, 1, 1, 1, 1, 1, 1},
					ExpectedResult: 800.0,
				},
				{
					Products:       []int{4},
					Amount:         []int{10},
					ExpectedResult: 4000.0,
				},
				{
					Products:       []int{7, 3},
					Amount:         []int{10, 10},
					ExpectedResult: 4000.0,
				},
				{
					Products:       []int{ 8},
					Amount:         []int{3},
					ExpectedResult: 4.5,
				},
				{
					Products:       []int{2, 3, 5},
					Amount:         []int{5, 3, 2},
					ExpectedResult: 2500.0,
				},
			},
		},
	}

	for idx, testcase := range testcases {
		constructorParam := testcase.Constructor
		cashier := Constructor(
			constructorParam.DiscountCheckpoint,
			constructorParam.Discount,
			constructorParam.Products,
			constructorParam.Prices,
		)
		for i, operation := range testcase.Operations {
			result := cashier.GetBill(operation.Products, operation.Amount)
			assert.Equalf(t, operation.ExpectedResult, result, "testcase: %d\noperation: %d", idx, i)
		}
	}
}
