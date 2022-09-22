package calc_test

import (
	"math/big"
	"testing"
)

func fibonacci(n int) int16 {
	dp := [1001]int16{0, 1}

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

func BenchmarkFibonacci(b *testing.B) {
	b.N = 100
	for i := 0; i < b.N; i++ {
		fibonacci(1000)
	}

}

func factorial(n int) *big.Int {
	var fact big.Int
	return fact.MulRange(1, int64(n))
}

func BenchmarkFactorial(b *testing.B) {
	b.N = 100
	for i := 0; i < b.N; i++ {
		factorial(100)
	}
}
