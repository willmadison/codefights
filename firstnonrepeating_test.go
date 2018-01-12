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
		{"bcbcbca", "a"},
		{"vbijvdpmxfztmlgieywuloeaudyokfjcoriqfwxuwdfxrllddihadvaeohgkjxiepvzmzhmpnuvgchqgabimpekppnewthrrbpvtfc", "_"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("firstNotRepeatingCharacter(%v)", tc.given), func(t *testing.T) {
			actual := firstNotRepeatingCharacter(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
