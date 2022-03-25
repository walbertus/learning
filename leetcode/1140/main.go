package leetcode

// https://leetcode.com/problems/stone-game-ii/
func stoneGameII(stonePiles []int) int {
    initialLimit := 1
    engine := newStoneGameIIEngine(stonePiles)
    result := engine.minMaxStoneTaken(0, initialLimit)
    return result
}

type stoneGameIIEngine struct {
    stonePiles []int
    memo       *memo
}

func (e *stoneGameIIEngine) minMaxStoneTaken(leftmostIndexAvailable, limit int) int {
    if leftmostIndexAvailable > e.getPilesLimit() {
        return 0
    }
    if e.getPilesLimit()-leftmostIndexAvailable < 2*limit {
        result := e.getStonesInPile(leftmostIndexAvailable, e.getPilesLimit())
        e.memo.setValue(leftmostIndexAvailable, limit, result)
        return result
    }
    if result, ok := e.memo.getValue(leftmostIndexAvailable, limit); ok {
        return result
    }
    maxStoneTaken := 0
    stoneLeft := e.getStonesInPile(leftmostIndexAvailable, e.getPilesLimit())
    for i := 0; i < limit*2; i++ {
        takenUpToIndex := leftmostIndexAvailable + i
        if takenUpToIndex > e.getPilesLimit() {
            break
        }
        newLimit := max(i+1, limit)
        stoneStolen := e.minMaxStoneTaken(takenUpToIndex+1, newLimit)
        stoneTaken := stoneLeft - stoneStolen
        maxStoneTaken = max(maxStoneTaken, stoneTaken)
    }
    e.memo.setValue(leftmostIndexAvailable, limit, maxStoneTaken)
    return maxStoneTaken
}

func (e *stoneGameIIEngine) getPilesLimit() int {
    return len(e.stonePiles) - 1
}

func (e *stoneGameIIEngine) getStonesInPile(from, to int) int {
    total := 0
    for i := from; i <= to; i++ {
        total += e.stonePiles[i]
    }
    return total
}

func newStoneGameIIEngine(stonePiles []int) *stoneGameIIEngine {
    return &stoneGameIIEngine{
        stonePiles: stonePiles,
        memo:       newMemo(len(stonePiles)),
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

type memo struct {
    containerArray [][]int
}

func (m *memo) setValue(key1, key2, value int) {
    m.containerArray[key1][key2] = value
}

func (m *memo) getValue(key1, key2 int) (int, bool) {
    if m.containerArray[key1][key2] == 0 {
        return 0, false
    }
    return m.containerArray[key1][key2], true
}

func newMemo(size int) *memo {
    containerArray := make([][]int, size+1)
    for i := range containerArray {
        containerArray[i] = make([]int, (size*2)+1)
        for j := range containerArray[i] {
            containerArray[i][j] = 0
        }
    }
    return &memo{
        containerArray: containerArray,
    }
}
