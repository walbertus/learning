package leetcode

// https://leetcode.com/problems/count-primes/
func countPrimes(n int) int {
    isNotPrime := make([]bool, n)
    for i := 2; i*i < n; i++ {
        if !isNotPrime[i] {
            for j := i * i; j < n; j += i {
                isNotPrime[j] = true
            }
        }
    }
    count := 0
    for i := 2; i < n; i++ {
        if !isNotPrime[i] {
            count++
        }
    }
    return count
}
