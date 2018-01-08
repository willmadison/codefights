package codefights

import "testing"

func TestReverseParenthesis(t *testing.T) {
	cases := []struct {
		given, expected string
	}{
		{
			"a(bc)de",
			"acbde",
		},
		{
			"a(bc)de",
			"acbde",
		},
	}

	for _, tc := range cases {
		t.Run(tc.given, func(t *testing.T) {
			actual := reverseParenthesis(tc.given)

			if actual != tc.expected {
				t.Error("got:", actual, "want:", tc.expected)
			}
		})
	}
}
