package leetcode

type Color int

const (
    white = Color(0)
    red   = Color(1)
    blue  = Color(2)
)

// https://leetcode.com/problems/is-graph-bipartite/
func isBipartite(graph [][]int) bool {
    disjointSet := newDisjointSet(len(graph))
    for id, adjacents := range graph {
        for _, adjacentID := range adjacents {
            disjointSet.union(id, adjacentID)
        }
    }
    startingNodes := map[int]bool{}
    for i := 0; i < len(graph); i++ {
        startingNodes[disjointSet.findParent(i)] = true
    }
    graphColor := make([]Color, len(graph))
    for node, _ := range startingNodes {
        if !canBipartiteFromNode(node, graph, graphColor) {
            return false
        }
    }
    return true
}

func canBipartiteFromNode(startingNode int, graph [][]int, graphColor []Color) bool {
    graphColor[startingNode] = red
    queue := []int{startingNode}
    for len(queue) != 0 {
        currentNode := queue[0]
        queue = queue[1:]

        for _, adjacent := range graph[currentNode] {
            if graphColor[adjacent] != white {
                if graphColor[adjacent] == graphColor[currentNode] {
                    return false
                }
                continue
            }
            graphColor[adjacent] = reverseColor(graphColor[currentNode])
            queue = append(queue, adjacent)
        }
    }
    return true
}

func reverseColor(color Color) Color {
    if color == white {
        panic("should not used in white color")
    }
    if color == red {
        return blue
    }
    return red
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
