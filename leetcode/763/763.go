package leetcode

import "sort"

const (
	OPEN  = iota
	CLOSE = iota
)

// https://leetcode.com/problems/partition-labels/submissions/
func partitionLabels(s string) []int {
	firstOccurrence := map[rune]int{}
	lastOccurrence := map[rune]int{}
	for idx, char := range s {
		if _, ok := firstOccurrence[char]; !ok {
			firstOccurrence[char] = idx
		}
		lastOccurrence[char] = idx
	}

	var actions []actionInLocation
	for r, i := range firstOccurrence {
		if firstOccurrence[r] == lastOccurrence[r] {
			continue
		}
		actions = append(actions, actionInLocation{
			action:       OPEN,
			charInEffect: r,
			location:     i,
		})
	}
	for r, i := range lastOccurrence {
		actions = append(actions, actionInLocation{
			action:       CLOSE,
			charInEffect: r,
			location:     i,
		})
	}

	sort.Sort(ByLocation(actions))
	var partitions []int
	set := newRuneSet()
	lastPartitionStart := 0
	for _, action := range actions {
		if action.action == OPEN {
			set.Add(action.charInEffect)
			continue
		}
		set.Pop(action.charInEffect)
		if set.Size() == 0 {
			partitions = append(partitions, action.location-lastPartitionStart+1)
			lastPartitionStart = action.location + 1
		}
	}
	return partitions
}

type actionCode int

type actionInLocation struct {
	action       actionCode
	charInEffect rune
	location     int
}

type ByLocation []actionInLocation

func (l ByLocation) Len() int           { return len(l) }
func (l ByLocation) Less(i, j int) bool { return l[i].location < l[j].location }
func (l ByLocation) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

type runeSet struct {
	container map[rune]bool
}

func newRuneSet() *runeSet {
	return &runeSet{map[rune]bool{}}
}

func (s *runeSet) Pop(r rune) {
	delete(s.container, r)
}

func (s *runeSet) Add(r rune) {
	s.container[r] = true
}

func (s *runeSet) Size() int {
	return len(s.container)
}
