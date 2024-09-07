package recursion

// find the maximum value
func FindMaxValue(dataset []int) int {
	// return if dataset has single value
	if len(dataset) == 1 {
		return dataset[0]
	}

	val1 := dataset[0]
	val2 := FindMaxValue(dataset[1:])
	if val1 > val2 {
		return val1
	} else {
		return val2
	}
}
