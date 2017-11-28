// TODO: document
package poethe_test

import (
	"image/color"
	"testing"

	"github.com/idagio/poethe/poethe"
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
		got, _ := generator.Generate(test.input)
		if len(got) != test.want {
			t.Errorf("Wanted (%d), got (%d)", test.want, len(got))
		}
	}
}
