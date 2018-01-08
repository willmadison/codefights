package codefights

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeRearranging(t *testing.T) {
	cases := []struct {
		given    string
		expected bool
	}{
		{
			"aabb",
			true,
		},
		{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabc",
			false,
		},
		{
			"abbcabb",
			true,
		},
		{
			"zyyzzzzz",
			true,
		},
		{
			"z",
			true,
		},
		{
			"zaa",
			true,
		},
		{
			"abca",
			false,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.given), func(t *testing.T) {
			actual := palindromeRearrange(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
