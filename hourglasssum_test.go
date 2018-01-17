package codefights

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHourGlassSum(t *testing.T) {
	cases := []struct {
		given    [][]int
		expected int
	}{
		{
			[][]int{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			},
			19,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("hourglassSum(%v)", tc.given), func(t *testing.T) {
			actual := hourglassSum(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
