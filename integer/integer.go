package integer

func Add(num1, num2 int) int {
	return num1 + num2
}

func Sum(arr []int) int {
	var sum int
	for _, numbers := range arr {
		sum += numbers
	}

	return sum
}

func SumAll(numbersToSum ...[]int) int {
	var sum int

	for _, numbers := range numbersToSum {
		for _, number := range numbers {
			sum += number
		}
	}

	return sum
}
