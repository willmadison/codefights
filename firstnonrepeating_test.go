package codefights

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstNotRepeatingCharacter(t *testing.T) {
	cases := []struct {
		given, expected string
	}{
		{"abacabad", "c"},
		{"abacabaabacaba", "_"},
		{"z", "z"},
		{"bcb", "c"},
		{"bcccccccb", "_"},
		{"abcdefghijklmnopqrstuvwxyziflskecznslkjfabe", "d"},
		{"zzz", "_"},
		{"bcccccccccccccyb", "y"},
		{"xdnxxlvupzuwgigeqjggosgljuhliybkjpibyatofcjbfxwtalc", "d"},
		{"ngrhhqbhnsipkcoqjyviikvxbxyphsnjpdxkhtadltsuxbfbrkof", "g"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("firstNotRepeatingCharacter(%v)", tc.given), func(t *testing.T) {
			actual := firstNotRepeatingCharacter(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestBitMasking(t *testing.T) {
	var bitMask int

	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("setBit(%v)/isBitSet(%v)", i, i), func(t *testing.T) {
			if i%2 == 0 {
				setBit(&bitMask, i)
			}
			assert.Equal(t, i%2 == 0, isBitSet(bitMask, i))
		})
	}
}
