package leetcode

import (
	"math"
	"strings"
)

// https://leetcode.com/problems/smallest-string-with-a-given-numeric-value/
func getSmallestString(n int, k int) string {
	valueLeftover := k - n
	numberOfCharNotA := int(math.Ceil(float64(valueLeftover) / 25.0))
	result := strings.Builder{}
	for i := 0; i < n-numberOfCharNotA; i++ {
		result.WriteRune('a')
	}
	if numberOfCharNotA != 0 && valueLeftover%25 != 0 {
		result.WriteRune(rune('a' + valueLeftover%25))
		numberOfCharNotA--
	}
	for i := 0; i < numberOfCharNotA; i++ {
		result.WriteRune('z')
	}
	return result.String()
}
