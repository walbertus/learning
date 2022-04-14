package leetcode

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type Direction int

const (
    Left Direction = iota
    Right
)

func longestZigZag(root *TreeNode) int {
    dp := make(memo)
    return traverseNode(root, dp) - 1
}

func traverseNode(node *TreeNode, dp memo) int {
    if node == nil {
        return 0
    }
    result := 1
    result = max(result, longestZigZagHelper(node, Left, dp))
    result = max(result, longestZigZagHelper(node, Right, dp))
    result = max(result, traverseNode(node.Left, dp))
    result = max(result, traverseNode(node.Right, dp))
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

type memo map[*TreeNode]map[Direction]int

func (m memo) get(key1 *TreeNode, key2 Direction) (int, bool) {
    if _, ok := m[key1]; !ok {
        return 0, false
    }
    val, ok := m[key1][key2]
    return val, ok
}

func (m memo) set(key1 *TreeNode, key2 Direction, value int) {
    if _, ok := m[key1]; !ok {
        m[key1] = make(map[Direction]int, 2)
    }
    m[key1][key2] = value
}

func longestZigZagHelper(node *TreeNode, direction Direction, dp memo) int {
    if node == nil {
        return 0
    }
    if val, ok := dp.get(node, direction); ok {
        return val
    }
    result := 1
    if direction == Left {
        result += longestZigZagHelper(node.Right, Right, dp)
    } else {
        result += longestZigZagHelper(node.Left, Left, dp)
    }
    dp.set(node, direction, result)
    return result
}
