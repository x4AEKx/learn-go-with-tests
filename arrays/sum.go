package main

func Sum(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var newSlice []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			newSlice = append(newSlice, 0)
		} else {
			tail := numbers[1:]
			newSlice = append(newSlice, Sum(tail))
		}

	}

	return newSlice
}
