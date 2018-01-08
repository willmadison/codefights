package codefights

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIPv4Address(t *testing.T) {
	cases := []struct {
		given    string
		expected bool
	}{
		{
			"172.16.254.1",
			true,
		},
		{
			"172.316.254.1",
			false,
		},
		{
			".254.255.0",
			false,
		},
		{
			"1.1.1.1a",
			false,
		},
		{
			"1",
			false,
		},
		{
			"0.254.255.0",
			true,
		},
		{
			"1.23.256.255.",
			false,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.given), func(t *testing.T) {
			actual := isIPv4Address(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
