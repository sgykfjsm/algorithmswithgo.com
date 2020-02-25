package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	var ret int
	for i := range numbers {
		ret += numbers[i]
	}

	return ret
}
