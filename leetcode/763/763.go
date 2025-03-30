package leetcode

// https://leetcode.com/problems/partition-labels/submissions/
func partitionLabels(s string) []int {
	firstOccurrence := newOccuranceMarker()
	lastOccurrence := newOccuranceMarker()
	stringLength := len(s)
	for idx, char := range s {
		if !firstOccurrence.isExist(char) {
			firstOccurrence.set(char, idx)
		}
		lastOccurrence.set(char, idx)
	}

	occuranceCounter := make([]int, stringLength)
	for i := 0; i < 26; i++ {
		currentRune := indexToRune(i)
		if firstOccurrence.isExist(currentRune) {
			occuranceCounter[firstOccurrence.get(currentRune)]++
		}
		if lastOccurrence.isExist(currentRune) {
			occuranceCounter[lastOccurrence.get(currentRune)]--
		}
	}

	partitions := make([]int, 0)
	counter := 0
	lastPartition := -1
	for i := 0; i < stringLength; i++ {
		counter += occuranceCounter[i]
		if counter == 0 {
			partitions = append(partitions, i-lastPartition)
			lastPartition = i
		}
	}

	return partitions
}

func runeToIndex(r rune) int {
	return int(r - 'a')
}

func indexToRune(i int) rune {
	return rune(i + 'a')
}

type occuranceMarker struct {
	container []int
}

func newOccuranceMarker() *occuranceMarker {
	container := make([]int, 26)
	for i := 0; i < 26; i++ {
		container[i] = -1
	}
	return &occuranceMarker{
		container: container,
	}
}

func (o *occuranceMarker) isExist(r rune) bool {
	return o.container[runeToIndex(r)] != -1
}

func (o *occuranceMarker) set(r rune, val int) {
	o.container[runeToIndex(r)] = val
}

func (o *occuranceMarker) get(r rune) int {
	return o.container[runeToIndex(r)]
}
