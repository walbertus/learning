package leetcode

import "errors"

func stoneGame(piles []int) bool {
	gameEngine := newStoneGameEngine()
	endGameDiff := gameEngine.takeStone(0, len(piles)-1, piles)
	if endGameDiff > 0 {
		return true
	} else {
		return false
	}
}

type stoneGameEngine struct {
	memo map[int]map[int]int
}

func newStoneGameEngine() *stoneGameEngine {
	return &stoneGameEngine{memo: map[int]map[int]int{}}
}

func (e *stoneGameEngine) takeStone(leftLimit, rightLimit int, piles []int) int {
	if leftLimit == rightLimit {
		return -piles[leftLimit]
	}
	result, err := e.getMemo(leftLimit, rightLimit)
	if err == nil {
		return result
	}
	if isPlayerOne(leftLimit, rightLimit) {
		takeLeft := e.takeStone(leftLimit+1, rightLimit, piles) + piles[leftLimit]
		takeRight := e.takeStone(leftLimit, rightLimit-1, piles) + piles[rightLimit]
		e.setMemo(leftLimit, rightLimit, max(takeLeft, takeRight))
		return max(takeLeft, takeRight)
	} else {
		takeLeft := e.takeStone(leftLimit+1, rightLimit, piles) - piles[leftLimit]
		takeRight := e.takeStone(leftLimit, rightLimit-1, piles) - piles[rightLimit]
		e.setMemo(leftLimit, rightLimit, min(takeLeft, takeRight))
		return min(takeLeft, takeRight)
	}
}

func (e *stoneGameEngine) setMemo(left, right, value int) {
	if _, ok := e.memo[left]; !ok {
		e.memo[left] = map[int]int{}
	}
	e.memo[left][right] = value
}

func (e *stoneGameEngine) getMemo(left, right int) (int, error) {
	if _, ok := e.memo[left]; !ok {
		return 0, errors.New("memo not found")
	}
	if _, ok := e.memo[left][right]; !ok {
		return 0, errors.New("memo not found")
	}
	return e.memo[left][right], nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isPlayerOne(left, right int) bool {
	return (right-left)%2 != 0
}
