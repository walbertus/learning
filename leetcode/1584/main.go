package leetcode

import "sort"

type Edge struct {
    source      int
    destination int
    cost        int
}

// https://leetcode.com/problems/min-cost-to-connect-all-points/
func minCostConnectPoints(points [][]int) int {
    numberOfPoints := len(points)
    if numberOfPoints == 1 {
        return 0
    }
    edgeList := generateSorterEdgeList(points)
    disjointSet := newDisjointSet(numberOfPoints)
    edgeTaken := 0
    result := 0
    currentIdx := 0
    for edgeTaken < numberOfPoints-1 {
        edge := edgeList[currentIdx]
        currentIdx++

        if disjointSet.isConnected(edge.source, edge.destination) {
            continue
        }

        result += edge.cost
        edgeTaken++
        disjointSet.union(edge.source, edge.destination)
    }
    return result
}

type DisjointSet struct {
    parents []int
}

func (s *DisjointSet) union(a, b int) {
    s.parents[s.findParent(a)] = s.findParent(b)
}

func (s *DisjointSet) isConnected(a, b int) bool {
    return s.findParent(a) == s.findParent(b)
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

func generateSorterEdgeList(points [][]int) []Edge {
    edgeList := make([]Edge, (len(points)*(len(points)-1))/2)
    idx := 0
    for i := 0; i < len(points); i++ {
        for j := i + 1; j < len(points); j++ {
            edgeList[idx] = Edge{
                source:      i,
                destination: j,
                cost:        calculateManhattanDistance(points[i], points[j]),
            }
            idx++
        }
    }
    sort.Slice(edgeList, func(i, j int) bool {
        return edgeList[i].cost < edgeList[j].cost
    })
    return edgeList
}

func calculateManhattanDistance(source, destination []int) int {
    return abs(source[0]-destination[0]) + abs(source[1]-destination[1])
}

func abs(val int) int {
    if val < 0 {
        return -val
    }
    return val
}
