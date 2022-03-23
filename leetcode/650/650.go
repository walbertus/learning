package leetcode

const maxPossibleSteps = 1000000

// https://leetcode.com/problems/2-keys-keyboard/
func minSteps(target int) int {
	if target == 1 {
		return 0
	}
	counter := newCounter(false)
	return counter.numberOfSteps(1, 1, target) + 1
}

type counter struct {
	memo       *memo
	enableMemo bool
}

func (c *counter) numberOfSteps(numberOfCharacter, charactersInClipboard, target int) int {
	if numberOfCharacter == target {
		return 0
	}
	if numberOfCharacter > target {
		return maxPossibleSteps
	}
	if c.enableMemo {
		if value, ok := c.memo.getValue(numberOfCharacter, charactersInClipboard); ok {
			return value
		}
	}
	stepsWhenPaste := c.numberOfSteps(numberOfCharacter+charactersInClipboard, charactersInClipboard, target) + 1
	stepsWhenCopy := c.numberOfSteps(numberOfCharacter+numberOfCharacter, numberOfCharacter, target) + 2
	result := min(stepsWhenCopy, stepsWhenPaste)
	if c.enableMemo {
		c.memo.setValue(numberOfCharacter, charactersInClipboard, result)
	}
	return result
}

func newCounter(enableMemo bool) *counter {
	return &counter{
		memo:       newMemo(),
		enableMemo: enableMemo,
	}
}

type memo struct {
	container map[int]map[int]int
}

func (m *memo) setValue(key1, key2, value int) {
	if _, ok := m.container[key1]; !ok {
		m.container[key1] = map[int]int{}
	}
	m.container[key1][key2] = value
}

func (m *memo) getValue(key1, key2 int) (int, bool) {
	if _, ok := m.container[key1]; !ok {
		return 0, false
	}
	value, ok := m.container[key1][key2]
	if !ok {
		return 0, false
	}
	return value, true
}

func newMemo() *memo {
	return &memo{container: map[int]map[int]int{}}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
