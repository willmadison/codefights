package codefights

func alternatingSums(a []int) []int {
	groups := []int{0, 0}

	for i := range a {
		groups[i%2] += a[i]
	}
	return groups
}
