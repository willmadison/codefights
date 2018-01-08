package codefights

func palindromeRearrange(inputString string) bool {
	characterCounts := map[rune]int{}

	for _, char := range inputString {
		characterCounts[char]++
	}

	var numOddCounts int
	for _, count := range characterCounts {
		if count%2 != 0 {
			numOddCounts++
		}
	}

	return numOddCounts < 2
}
