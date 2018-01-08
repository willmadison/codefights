package codefights

func areEquallyStrong(yourLeft, yourRight, friendsLeft, friendsRight int) bool {
	yourMin, yourMax := minMax(yourLeft, yourRight)
	friendsMin, friendsMax := minMax(friendsLeft, friendsRight)
	return yourMin == friendsMin && yourMax == friendsMax
}

func minMax(a, b int) (int, int) {
	var min, max int

	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	return min, max
}
