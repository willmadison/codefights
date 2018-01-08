package codefights

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreSimilar(t *testing.T) {
	cases := []struct {
		given struct {
			a, b []int
		}
		expected bool
	}{
		{
			struct {
				a, b []int
			}{
				a: []int{1, 2, 3},
				b: []int{1, 2, 3},
			},
			true,
		},
		{
			struct {
				a, b []int
			}{
				a: []int{1, 2, 3},
				b: []int{2, 1, 3},
			},
			true,
		},
		{
			struct {
				a, b []int
			}{
				a: []int{1, 2, 2},
				b: []int{2, 1, 1},
			},
			false,
		},
		{
			struct {
				a, b []int
			}{
				a: []int{1, 1, 4},
				b: []int{1, 2, 3},
			},
			false,
		},
		{
			struct {
				a, b []int
			}{
				a: []int{1, 2, 3},
				b: []int{1, 10, 2},
			},
			false,
		},
		{
			struct {
				a, b []int
			}{
				a: []int{2, 3, 1},
				b: []int{1, 3, 2},
			},
			true,
		},
		{
			struct {
				a, b []int
			}{
				a: []int{2, 3, 9},
				b: []int{10, 3, 2},
			},
			false,
		},
		{
			struct {
				a, b []int
			}{
				a: []int{4, 6, 3},
				b: []int{3, 4, 6},
			},
			false,
		},
		{
			struct {
				a, b []int
			}{
				a: []int{832, 998, 148, 570, 533, 561, 894, 147, 455, 279},
				b: []int{832, 998, 148, 570, 533, 561, 455, 147, 894, 279},
			},
			true,
		},
		{
			struct {
				a, b []int
			}{
				a: []int{832, 998, 148, 570, 533, 561, 894, 147, 455, 279},
				b: []int{832, 570, 148, 998, 533, 561, 455, 147, 894, 279},
			},
			false,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.given), func(t *testing.T) {
			actual := areSimilar(tc.given.a, tc.given.b)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
