package codefights

func firstNotRepeatingCharacter(s string) string {
	var bitMask int

	lastUniqueSeen, previousLastUnique := '_', '_'

	for index := len(s) - 1; index >= 0; index-- {
		c := rune(s[index])
		if seen := isBitSet(bitMask, int(c-'a')); !seen {
			lastUniqueSeen, previousLastUnique = c, lastUniqueSeen
			setBit(&bitMask, int(c-'a'))
		} else if lastUniqueSeen == c {
			lastUniqueSeen = previousLastUnique
		}
	}

	return string(lastUniqueSeen)
}

func setBit(given *int, i int) {
	*given |= (1 << uint(i))
}

func isBitSet(given int, i int) bool {
	return given&(1<<uint(i)) != 0
}
