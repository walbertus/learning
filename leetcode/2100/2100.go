package leetcode


// https://leetcode.com/problems/find-good-days-to-rob-the-bank/
func goodDaysToRobBank(securities []int, safeTime int) []int {
	daysBeforeWithNonIncreasingSecurity := computeDaysBeforeWithNonIncreasingSecurity(securities)
	daysAfterWithNonDecreasingSecurity := computeDaysAfterWithNonDecreasingSecurity(securities)

	goodDays := []int{}
	for i := 0; i < len(securities); i++ {
		if daysBeforeWithNonIncreasingSecurity[i] >= safeTime && daysAfterWithNonDecreasingSecurity[i] >= safeTime {
			goodDays = append(goodDays, i)
		}
	}
	return goodDays
}

func computeDaysBeforeWithNonIncreasingSecurity(securities []int) []int {
	daysBefore := []int{0}
	for i := 1; i < len(securities); i++ {
		if securities[i] <= securities[i-1] {
			daysBefore = append(daysBefore, daysBefore[i-1]+1)
		} else {
			daysBefore = append(daysBefore, 0)
		}
	}
	return daysBefore
}

func computeDaysAfterWithNonDecreasingSecurity(securities []int) []int {
	daysAfter := []int{0}
	for i := 1; i < len(securities); i++ {
		daysAfter = append(daysAfter, 0)
	}
	for i := len(securities) - 2; i >= 0; i-- {
		if securities[i] <= securities[i+1] {
			daysAfter[i] = daysAfter[i+1] + 1
		} else {
			daysAfter[i] = 0
		}
	}
	return daysAfter
}
