package main

func Sum(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	numbersToSumLength := len(numbersToSum)
	newSlice := make([]int, numbersToSumLength)

	for i := 0; i < numbersToSumLength; i++ {
		newSlice[i] = Sum(numbersToSum[i])
	}

	return newSlice
}
