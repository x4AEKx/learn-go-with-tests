package main

func Sum(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var newSlice []int

	for _, value := range numbersToSum {
		newSlice = append(newSlice, Sum(value))
	}

	return newSlice
}
