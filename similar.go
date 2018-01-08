package codefights

func areSimilar(a, b []int) bool {
	var differences []int

	for i := range a {
		if a[i] != b[i] {
			differences = append(differences, i)
		}
	}

	if len(differences) > 2 || len(differences) == 1 {
		return false
	}

	if len(differences) == 0 {
		return true
	}

	return a[differences[0]] == b[differences[1]] && a[differences[1]] == b[differences[0]]
}
