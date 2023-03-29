package functions

import (
	"math/rand"
)

func NumGenerator(numCount int, maxNumber int) []int {

	numbers := make(map[int]bool)

	for len(numbers) < numCount {
		n := rand.Intn(maxNumber)
		if !numbers[n] {
			numbers[n] = true
		}
	}

	var result []int
	for n := range numbers {
		result = append(result, n)
	}

	return result
}
