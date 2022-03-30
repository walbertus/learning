package leetcode

// https://leetcode.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	suspectedRow := binarySearchRow(matrix, target)
	return binarySearchColumn(matrix[suspectedRow], target)
}

func binarySearchRow(matrix [][]int, target int) (rowNumber int) {
	leftPointer := 0
	rightPointer := len(matrix)-1
	for leftPointer < rightPointer {
		midPointer := (leftPointer + rightPointer + 1) >> 1
		if matrix[midPointer][0] > target {
			rightPointer = midPointer-1
		} else {
			leftPointer = midPointer
		}
	}
	return leftPointer
}

func binarySearchColumn(row []int, target int) bool {
	leftPointer := 0
	rightPointer := len(row)-1
	for leftPointer <= rightPointer {
		midPointer := (leftPointer + rightPointer) >> 1
		if row[midPointer] == target {
			return true
		}
		if row[midPointer] > target {
			rightPointer = midPointer - 1
		} else {
			leftPointer = midPointer + 1
		}
	}
	return false
}
