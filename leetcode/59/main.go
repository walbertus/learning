package leetcode

// https://leetcode.com/problems/spiral-matrix-ii/
func generateMatrix(n int) [][]int {
	var walker MatrixWalker
	walker = newMatrixWalker(n)
	for i := 0; i < n*n; i++ {
		walker.walk()
	}
	return walker.result()
}

type MatrixWalker interface {
	walk()
	result() [][]int
}

type matrixWalker struct {
	matrix [][]int
	walkerState walkFunc
	currentRow, currentColumn, counter int
}

func (w *matrixWalker) walk() {
	w.walkerState(w)
}

func (w *matrixWalker) rowSize() int {
	return len(w.matrix)
}

func (w *matrixWalker) columnSize() int {
	return len(w.matrix[0])
}

func (w *matrixWalker) isCellEmpty(row, column int) bool {
	if row < 0 || row >= w.rowSize() || column < 0 || column >= w.columnSize() {
		return false
	}
	return w.matrix[row][column] == 0
}

func (w *matrixWalker) fillCurrentCell() {
	w.counter++
	w.matrix[w.currentRow][w.currentColumn] = w.counter
}

func (w *matrixWalker) result() [][]int {
	return w.matrix
}

func newMatrixWalker(n int) *matrixWalker {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	return &matrixWalker{
		matrix:      matrix,
		walkerState: walkRightFunc,
		currentRow: 0,
		currentColumn: -1,
		counter: 0,
	}
}

type walkFunc func(parent *matrixWalker)

func walkRightFunc(parent *matrixWalker) {
	if !parent.isCellEmpty(parent.currentRow, parent.currentColumn+1) {
		parent.walkerState = walkDownFunc
		parent.walk()
		return
	}
	parent.currentColumn++
	parent.fillCurrentCell()
}

func walkDownFunc(parent *matrixWalker) {
	if !parent.isCellEmpty(parent.currentRow+1, parent.currentColumn) {
		parent.walkerState = walkLeftFunc
		parent.walk()
		return
	}
	parent.currentRow++
	parent.fillCurrentCell()
}

func walkLeftFunc(parent *matrixWalker) {
	if !parent.isCellEmpty(parent.currentRow, parent.currentColumn-1) {
		parent.walkerState = walkUpFunc
		parent.walk()
		return
	}
	parent.currentColumn--
	parent.fillCurrentCell()
}

func walkUpFunc(parent *matrixWalker) {
	if !parent.isCellEmpty(parent.currentRow-1, parent.currentColumn) {
		parent.walkerState = walkRightFunc
		parent.walk()
		return
	}
	parent.currentRow--
	parent.fillCurrentCell()
}
