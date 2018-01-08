package codefights

func arrayMaximalAdjacentDifference(values []int) int {
	var maxDiff int

	for i := 0; i < len(values)-1; i++ {
		diff := abs(values[i] - values[i+1])

		if diff > maxDiff {
			maxDiff = diff
		}
	}

	return maxDiff
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
