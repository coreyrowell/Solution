package utils

func contains(numbers []int, number int) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}