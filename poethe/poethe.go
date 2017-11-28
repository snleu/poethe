// TODO: document
package poethe

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

// TODO: document
// TODO: *maybe& define type `Palette` struct
func ParseColors(input string) ([]*color.RGBA, error) {
	entries := strings.Split(input, ",")
	colors := make([]*color.RGBA, len(entries))

	for idx, entry := range entries {
		// First check if the entry has the correct format
		if !strings.HasPrefix(entry, "#") {
			return nil, fmt.Errorf("color not prefixed by '#': %s", entry)
		}

		vals := strings.SplitAfter(entry, "")
		hexRequiredLength := int(7)
		if len(vals) != hexRequiredLength {
			return nil, fmt.Errorf("color does not have length of %d: %s", hexRequiredLength, entry)
		}

		redString := strings.Join(vals[1:3], "")
		red, err := strconv.ParseUint(redString, 16, 8)
		if err != nil {
			return nil, fmt.Errorf("could not decode hex for 'red' color: %v", err)
		}

		greenString := strings.Join(vals[3:5], "")
		green, err := strconv.ParseUint(greenString, 16, 8)
		if err != nil {
			return nil, fmt.Errorf("could not decode hex for 'green' color: %v", err)
		}

		blueString := strings.Join(vals[5:7], "")
		blue, err := strconv.ParseUint(blueString, 16, 8)
		if err != nil {
			return nil, fmt.Errorf("could not decode hex for 'blue' color: %v", err)
		}

		rg := &color.RGBA{
			R: uint8(red),
			G: uint8(green),
			B: uint8(blue),
			A: 1,
		}

		colors[idx] = rg
	}

	return colors, nil
}

// TODO: document
type Generator interface {
	Description() string
	Generate(colors []*color.RGBA) []*color.RGBA
}

// TODO: document
func FormatColors(colors []*color.RGBA) string {
	colorStrings := make([]string, len(colors))
	for idx, val := range colors {
		red := doubleDigitHex(val.R)
		green := doubleDigitHex(val.G)
		blue := doubleDigitHex(val.B)
		colorStrings[idx] = fmt.Sprintf("#%s%s%s", red, green, blue)
	}

	return strings.Join(colorStrings, " ")
}

func doubleDigitHex(val uint8) string {
	if val <= 15 {
		return fmt.Sprintf("0%x", val)
	}
	return fmt.Sprintf("%x", val)
}
