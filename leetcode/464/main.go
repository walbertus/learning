package leetcode

// https://leetcode.com/problems/can-i-win/
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
    if desiredTotal <= maxChoosableInteger {
        return true
    }
    if maxChoosableInteger*(maxChoosableInteger+1) < desiredTotal*2 {
        return false
    }
    choosableInteger := bitmask(0)
    for i := 1; i <= maxChoosableInteger; i++ {
        choosableInteger = choosableInteger.add(i)
    }
    counter := newGameCounter()
    return counter.canIWin(choosableInteger, desiredTotal)
}

type memo map[bitmask]bool

type gameCounter struct {
    memo memo
}

func (c *gameCounter) canIWin(choosableInteger bitmask, leftover int) bool {
    if leftover <= 0 {
        return true
    }
    if value, ok := c.memo[choosableInteger]; ok {
        return value
    }

    choosableIntegers := choosableInteger.getInts()
    enemyWillNotLose := true
    for _, integer := range choosableIntegers {
        if integer >= leftover {
            enemyWillNotLose = false
            break
        }
        enemyWillNotLose = enemyWillNotLose && c.canIWin(choosableInteger.remove(integer), leftover-integer)
        if !enemyWillNotLose {
            break
        }
    }
    c.memo[choosableInteger] = !enemyWillNotLose
    return !enemyWillNotLose
}

func newGameCounter() *gameCounter {
    return &gameCounter{
        memo: memo{},
    }
}

type bitmask uint32

func (b bitmask) add(number int) bitmask {
    return b | (1 << number)
}

func (b bitmask) remove(number int) bitmask {
    return b & ^(1 << number)
}

func (b bitmask) getInts() []int {
    pointer := 1
    var ints []int
    for (1 << pointer) <= int(b) {
        if (b & (1 << pointer)) != 0 {
            ints = append(ints, pointer)
        }
        pointer++
    }
    return ints
}
