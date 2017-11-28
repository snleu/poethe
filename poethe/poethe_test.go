// TODO: document
package poethe_test

import (
	"github.com/idagio/poethe/poethe"
	"image/color"
	"reflect"
	"testing"
)

func TestParseColors(t *testing.T) {
	var tests = []struct {
		input string
		want  []*color.RGBA
	}{
		{"#000000", []*color.RGBA{
			{R: 0, G: 0, B: 0, A: 1},
		}},
		{"#FF0000", []*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
		}},
		{"#ff0000", []*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
		}},
		{"#000000,#ff0000", []*color.RGBA{
			{R: 0, G: 0, B: 0, A: 1},
			{R: 255, G: 0, B: 0, A: 1},
		}},
		{"#000000,#FF0000", []*color.RGBA{
			{R: 0, G: 0, B: 0, A: 1},
			{R: 255, G: 0, B: 0, A: 1},
		}},
		{"#ff0000,#00ff00", []*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
			{R: 0, G: 255, B: 0, A: 1},
		}},
		{"#ff0000,#00ff00,#0000ff", []*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
			{R: 0, G: 255, B: 0, A: 1},
			{R: 0, G: 0, B: 255, A: 1},
		}},
		{"#ff0000,#0000ff,#00ff00", []*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
			{R: 0, G: 0, B: 255, A: 1},
			{R: 0, G: 255, B: 0, A: 1},
		}},
	}

	for _, test := range tests {
		if got, _ := poethe.ParseColors(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Wanted (%s), got (%s)", test.want, got)
		}
	}
}

func TestFormatColors(t *testing.T) {
	var tests = []struct {
		input []*color.RGBA
		want  string
	}{
		{[]*color.RGBA{
			{R: 0, G: 0, B: 0, A: 1},
		}, "#000000"},
		{[]*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
		}, "#ff0000"},
		{[]*color.RGBA{
			{R: 107, G: 12, B: 228, A: 1},
		}, "#6b0ce4"},
		{[]*color.RGBA{
			{R: 0, G: 0, B: 0, A: 1},
			{R: 255, G: 0, B: 0, A: 1},
		}, "#000000 #ff0000"},
		{[]*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
			{R: 0, G: 255, B: 0, A: 1},
		}, "#ff0000 #00ff00"},
		{[]*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
			{R: 0, G: 255, B: 0, A: 1},
			{R: 0, G: 0, B: 255, A: 1},
		}, "#ff0000 #00ff00 #0000ff"},
		{[]*color.RGBA{
			{R: 255, G: 0, B: 0, A: 1},
			{R: 0, G: 0, B: 255, A: 1},
			{R: 0, G: 255, B: 0, A: 1},
		}, "#ff0000 #0000ff #00ff00"},
	}

	for _, test := range tests {
		if got := poethe.FormatColors(test.input); got != test.want {
			t.Errorf("Wanted (%s), got (%s)", test.want, got)
		}
	}
}
