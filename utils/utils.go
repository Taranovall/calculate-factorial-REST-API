package utils

import (
	"calculateFactorial/models"
	"sync"
)

func ValidateNumbers(factorialRequest models.FactorialRequest) bool {
	if factorialRequest.AFactorial == nil || *factorialRequest.AFactorial < 0 {
		return false
	}

	if factorialRequest.BFactorial == nil || *factorialRequest.BFactorial < 0 {
		return false
	}

	return true
}

func RunFactorialCalculations(a, b int) map[string]int {
	type Result struct {
		Index     string
		Factorial int
	}

	ch := make(chan Result, 2)
	results := make(map[string]int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		ch <- Result{"a!", calculateFactorial(a)}
	}()
	go func() {
		defer wg.Done()
		ch <- Result{"b!", calculateFactorial(b)}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		results[result.Index] = result.Factorial
	}

	return results
}

func calculateFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calculateFactorial(n-1)
}
