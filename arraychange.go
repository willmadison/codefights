package codefights

func arrayChange(inputArray []int) int {
	var moves int

	for i := 0; i < len(inputArray)-1; i++ {
		if inputArray[i] >= inputArray[i+1] {
			moves += inputArray[i] + 1 - inputArray[i+1]
			inputArray[i+1] = inputArray[i] + 1
		}
	}

	return moves
}
