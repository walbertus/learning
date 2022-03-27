package leetcode

import "sort"

const soldierSign = 1

type bySoldier []rowData

func (b bySoldier) Len() int {
    return len(b)
}

func (b bySoldier) Less(i, j int) bool {
    if b[i].soldier == b[j].soldier {
        return b[i].index < b[j].index
    }
    return b[i].soldier < b[j].soldier
}

func (b bySoldier) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}

type rowData struct {
    index   int
    soldier int
}

// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/
func kWeakestRows(matrix [][]int, query int) []int {
    rowsData := []rowData{}
    for id, row := range matrix {
        rowsData = append(rowsData, rowData{
            index:   id,
            soldier: numberOfSoldier(row),
        })
    }
    sort.Sort(bySoldier(rowsData))
    var result []int
    for i := 0; i < query; i++ {
        result = append(result, rowsData[i].index)
    }
    return result
}

func numberOfSoldier(list []int) int {
    leftPointer := 0
    rightPointer := len(list) - 1
    for leftPointer < rightPointer {
        midPointer := (leftPointer + rightPointer + 1) / 2
        if list[midPointer] == soldierSign {
            leftPointer = midPointer
        } else {
            rightPointer = midPointer - 1
        }
    }
    if list[leftPointer] != soldierSign {
        return 0
    }
    return leftPointer + 1
}
