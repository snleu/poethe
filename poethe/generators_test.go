// TODO: document
package poethe_test

import (
	"github.com/idagio/poethe/poethe"
	"image/color"
	"testing"
)

func TestRandomGenerator(t *testing.T) {
	generator := poethe.RandomGenerator{}

	var tests = []struct {
		input []*color.RGBA
		want  int
	}{
		{[]*color.RGBA{
			{R: 0, G: 0, B: 0, A: 1},
		}, 5},
		{[]*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
		}, 5},
		{[]*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
		}, 5},
		{[]*color.RGBA{
			{R: 0, G: 0, B: 0, A: 1},
			{R: 255, G: 0, B: 0, A: 1},
		}, 5},
		{[]*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
			{R: 0, G: 255, B: 0, A: 1},
		}, 5},
		{[]*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
			{R: 0, G: 255, B: 0, A: 1},
			{R: 0, G: 0, B: 255, A: 1},
		}, 5},
		{[]*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
			{R: 0, G: 0, B: 255, A: 1},
			{R: 0, G: 255, B: 0, A: 1},
		}, 5},
	}

	for _, test := range tests {
		if got := generator.Generate(test.input); len(got) != test.want {
			t.Errorf("Wanted (%d), got (%d)", test.want, len(got))
		}
	}
}
