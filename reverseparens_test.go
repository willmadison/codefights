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
			"a(bcdefghijkl(mno)p)q",
			"apmnolkjihgfedcbq",
		},
		{
			"co(de(fight)s)",
			"cosfighted",
		},
		{
			"Code(Cha(lle)nge)",
			"CodeegnlleahC",
		},
		{
			"Where are the parentheses?",
			"Where are the parentheses?",
		},
		{
			"abc(cba)ab(bac)c",
			"abcabcabcabc",
		},
		{
			"The ((quick (brown) (fox) jumps over the lazy) dog)",
			"The god quick nworb xof jumps over the lazy",
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
