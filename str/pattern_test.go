package str

import (
	"strings"
	"testing"
)

type testCase struct {
	input   string
	pattern string
	index   int
}

var (
	cases = []testCase{
		{
			input:   "JIM_SAW_ME_IN_A_BARBERSHOP",
			pattern: "BARBER",
			index:   -1,
		},
	}
)

func init() {
	for i, c := range cases {
		cases[i].index = strings.Index(c.input, c.pattern)
	}
}
func TestHorspoolSearch(t *testing.T) {
	for i, c := range cases {
		index := horspoolSearch(c.input, c.pattern)
		if index != c.index {
			t.Fatalf("test case:%d want:%d got:%d\n", i, c.index, index)
			return
		}
	}
}
