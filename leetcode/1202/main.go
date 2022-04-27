package leetcode

import (
    "sort"
    "unsafe"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
    disjointSet := newDisjointSet(len(s))
    for _, pair := range pairs {
        disjointSet.union(pair[0], pair[1])
    }
    connectedIndices := map[int][]int{}
    for i := 0; i < len(s); i++ {
        parent := disjointSet.findParent(i)
        if _, ok := connectedIndices[parent]; !ok {
            connectedIndices[parent] = []int{}
        }
        components := connectedIndices[parent]
        components = append(components, i)
        connectedIndices[parent] = components
    }
    resultInArr := make([]byte, len(s))
    for _, indices := range connectedIndices {
        sort.Ints(indices)
        chars := make([]byte, len(indices))
        idx := 0
        for _, index := range indices {
            chars[idx] = s[index]
            idx++
        }
        sort.Slice(chars, func(i, j int) bool {
            return chars[i] < chars[j]
        })
        idx = 0
        for _, index := range indices {
            resultInArr[index] = chars[idx]
            idx++
        }
    }
    return *(*string)(unsafe.Pointer(&resultInArr))
}

type DisjointSet struct {
    parents []int
}

func (s *DisjointSet) union(a, b int) {
    s.parents[s.findParent(a)] = s.findParent(b)
}

func (s *DisjointSet) findParent(id int) int {
    if s.parents[id] != id {
        s.parents[id] = s.findParent(s.parents[id])
    }
    return s.parents[id]
}

func newDisjointSet(size int) *DisjointSet {
    parents := make([]int, size)
    for i := 0; i < size; i++ {
        parents[i] = i
    }
    return &DisjointSet{parents: parents}
}
