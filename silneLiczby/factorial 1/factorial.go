package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func silnia(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	return big.NewInt(int64(n)).Mul(big.NewInt(int64(n)), silnia(n-1))
}

func fibonacci(n int, calls *[31]int) int {
	calls[n] += 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1, calls) + fibonacci(n-2, calls)
}

func isIn(str string, substr string) bool {
	if len(str) < len(substr) {
		return false
	}
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func findMyNumbers(ascii []int) bool {
	i := 1
	maxIterations := 10000

	for i <= maxIterations {
		silnia := silnia(i).String()

		if allCodesFound(silnia, ascii) {
			fmt.Println("Silna liczba: ", i)
			printWeakNumber(i)
			return true
		}

		i++
	}
	return false
}

func allCodesFound(digits string, ascii []int) bool {
	for _, code := range ascii {
		if !isIn(digits, strconv.Itoa(code)) {
			return false
		}
	}
	return true
}

func printWeakNumber(i int) {
	var calls [31]int
	fibonacci(30, &calls)
	var closest int
	for j := 0; j < len(calls); j++ {
		if math.Abs(float64(i-calls[j])) < math.Abs(float64(i-calls[closest])) {
			closest = j
		}
	}
	fmt.Println("Słaba liczba: ", closest)
}
