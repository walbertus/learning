package leetcode

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one
func numSteps(s string) int {
    if s == "1" {
        return 0
    }
    res := 0
    remainder := false
    for i := len(s) - 1; i > 0; i-- {
        ch := s[i]
        if ch == '1' {
            if remainder {
                res++
            } else {
                res += 2
                remainder = true
            }
        } else {
            if remainder {
                res += 2
            } else {
                res++
            }
        }
    }
    if remainder {
        res++
    }
    return res
}
