package leetcode

import (
	"math"
	"sync"
)

// https://leetcode.com/problems/stone-game-iv/
func winnerSquareGame(n int) bool {
	if n <= 0 {
		return false
	}
	if ok, val := getDPInstance().get(n); ok {
		return val
	}
	if isQuadratic(n) {
		getDPInstance().set(n, true)
		return true
	}
	isEnemyAlwaysWin := true
	for i := 1; i*i <= n; i++ {
		isEnemyAlwaysWin = isEnemyAlwaysWin && winnerSquareGame(n-i*i)
		if !isEnemyAlwaysWin {
			break
		}
	}
	getDPInstance().set(n, !isEnemyAlwaysWin)
	return !isEnemyAlwaysWin
}

func isQuadratic(val int) bool {
	if val == 0 {
		return true
	}
	root := int(math.Sqrt(float64(val)))
	return root*root == val
}

type dp struct {
	container   []bool
	hasComputed []bool
}

func (d *dp) get(n int) (bool, bool) {
	if d.hasComputed[n] {
		return true, d.container[n]
	}
	return false, false
}

func (d *dp) set(n int, val bool) {
	d.container[n] = val
	d.hasComputed[n] = true
}

func newDP(n int) *dp {
	return &dp{
		container:   make([]bool, n+1),
		hasComputed: make([]bool, n+1),
	}
}

var dpInstance *dp

var once sync.Once

func getDPInstance() *dp {
	if dpInstance == nil {
		once.Do(func() {
			dpInstance = newDP(100000)
		})
	}
	return dpInstance
}
