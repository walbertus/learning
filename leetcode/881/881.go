package leetcode

import "sort"

// https://leetcode.com/problems/boats-to-save-people
func numRescueBoats(peopleWeight []int, boatLimit int) int {
	sort.Ints(peopleWeight)
	numberOfBoats := 0
	leftPointer := 0
	rightPointer := len(peopleWeight) - 1
	for leftPointer <= rightPointer {
		boatCapacityLeft := boatLimit - peopleWeight[rightPointer]
		rightPointer--
		if boatCapacityLeft >= peopleWeight[leftPointer] {
			leftPointer++
		}
		numberOfBoats++
	}
	return numberOfBoats
}
