package main 

func Sum(numbers []int) int {
	
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int{
	var result []int
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return result
}