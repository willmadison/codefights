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
		{
			[][]int{
				{0, -4, -6, 0, -7, -6},
				{-1, -2, -6, -8, -3, -1},
				{-8, -4, -2, -8, -8, -6},
				{-3, -1, -2, -5, -7, -4},
				{-3, -5, -3, -6, -6, -6},
				{-3, -6, 0, -8, -6, -7},
			},
			-19,
		},
		{
			[][]int{
				{-1, -1, 0, -9, -2, -2},
				{-2, -1, -6, -8, -2, -5},
				{-1, -1, -1, -2, -3, -4},
				{-1, -9, -2, -4, -4, -5},
				{-7, -3, -3, -2, -9, -9},
				{-1, -3, -1, -2, -4, -5},
			},
			-6,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("hourglassSum(%v)", tc.given), func(t *testing.T) {
			actual := hourglassSum(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
