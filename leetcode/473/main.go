package leetcode

// https://leetcode.com/problems/matchsticks-to-square/
func makesquare(matchsticks []int) bool {
    if len(matchsticks) < 4 {
        return false
    }
    perimeter := getPerimeter(matchsticks)
    if perimeter%4 != 0 {
        return false
    }
    expectedSide := perimeter / 4
    matchsticksBitmask := bitmask(0)
    for i := 0; i < len(matchsticks); i++ {
        matchsticksBitmask = matchsticksBitmask.add(i)
    }
    dp := memo{}
    return canMakeSquare(expectedSide, 0, expectedSide, matchsticksBitmask, matchsticks, dp)
}

func canMakeSquare(currentSideNeed int, completedSide int, expectedSide int, bit bitmask, matchsticks []int, dp memo) bool {
    if completedSide == 3 {
        return true
    }
    if value, ok := dp[bit]; ok {
        return value
    }
    possible := false
    for _, stickID := range bit.getAvailableStickID() {
        currentStick := matchsticks[stickID]
        currentSideLeft := currentSideNeed-currentStick
        if currentSideLeft > 0 {
            possible = possible || canMakeSquare(currentSideLeft, completedSide, expectedSide, bit.use(stickID), matchsticks, dp)
        } else if currentSideLeft == 0 {
            possible = possible || canMakeSquare(expectedSide, completedSide+1, expectedSide, bit.use(stickID), matchsticks, dp)
        }
        if possible {
            break
        }
    }
    dp[bit] = possible
    return possible
}

func getPerimeter(matchsticks []int) int {
    perimeter := 0
    for _, matchstick := range matchsticks {
        perimeter += matchstick
    }
    return perimeter
}

type memo map[bitmask]bool

type bitmask uint16

func (b bitmask) getAvailableStickID() []int {
    pointer := 0
    var availableID []int
    comparator := bitmask(1 << pointer)
    for comparator <= b {
        if (comparator & b) != 0 {
            availableID = append(availableID, pointer)
        }
        pointer++
        comparator = bitmask(1 << pointer)
    }
    return availableID
}

func (b bitmask) add(id int) bitmask {
    return b | (1 << id)
}

func (b bitmask) use(id int) bitmask {
    return b & ^(1 << id)
}
