// TODO: document
package poethe_test

import (
	"github.com/snleu/poethe/poethe"
	"testing"
)

func TestParseColors(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"hello", "hello"},
	}

	for _, test := range tests {
		if got := poethe.HelloWorld(test.input); got == test.want {
			t.Errorf("Wanted (%s), got (%s)", test.want, got)
		}
	}
}
