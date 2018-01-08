package codefights

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlternatingSums(t *testing.T) {
	cases := []struct {
		given, expected []int
	}{
		{
			[]int{50, 60, 60, 45, 70},
			[]int{180, 105},
		},
		{
			[]int{100, 50},
			[]int{100, 50},
		},
		{
			[]int{80},
			[]int{80, 0},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.given), func(t *testing.T) {
			actual := alternatingSums(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
