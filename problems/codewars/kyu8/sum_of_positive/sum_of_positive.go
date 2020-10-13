package main

func PositiveSum(numbers []int) int {
	var sum int = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			sum += numbers[i]
		}
	}

	return sum
}
