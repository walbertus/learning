package leetcode

import "strings"

// https://leetcode.com/problems/n-queens/
func solveNQueens(n int) [][]string {
    if n == 1 {
        return [][]string{{"Q"}}
    }
    if n == 2 {
        return [][]string{}
    }
    if n == 3 {
        return [][]string{}
    }
    var board [][]int
    for i := 0; i < n; i++ {
        board = append(board, []int{})
        for j := 0; j < n; j++ {
            board[i] = append(board[i], 0)
        }
    }
    return findSolutions(0, board, n)
}

func findSolutions(row int, board [][]int, size int) [][]string {
    if row >= size {
        var stringBoard [][]string
        stringBoard = append(stringBoard, convertBoard(board, size))
        return stringBoard
    }
    result := [][]string{}
    for column := 0; column < size; column++ {
        if !isSafe(board, size, row, column) {
            continue
        }
        board[row][column] = 1
        result = append(result, findSolutions(row+1, board, size)...)
        board[row][column] = 0
    }
    return result
}

func isSafe(board [][]int, size int, row, column int) bool {
    for i := 0; i < size; i++ {
        if board[i][column] == 1 {
            return false
        }
        if board[row][i] == 1 {
            return false
        }
    }
    for i, j := row, column; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == 1 {
            return false
        }
    }
    for i, j := row, column; i < size && j < size; i, j = i+1, j+1 {
        if board[i][j] == 1 {
            return false
        }
    }
    for i, j := row, column; i >= 0 && j < size; i, j = i-1, j+1 {
        if board[i][j] == 1 {
            return false
        }
    }
    for i, j := row, column; i < size && j > 0; i, j = i+1, j-1 {
        if board[i][j] == 1 {
            return false
        }
    }
    return true
}

func convertBoard(board [][]int, size int) []string {
    var result []string
    for i := 0; i < size; i++ {
        str := strings.Builder{}
        for j := 0; j < size; j++ {
            if board[i][j] == 0 {
                str.WriteByte('.')
            } else {
                str.WriteByte('Q')
            }
        }
        result = append(result, str.String())
    }
    return result
}
