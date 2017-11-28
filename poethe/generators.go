// TODO: document
package poethe

import (
	"image/color"
	"math/rand"
)

const (
	numberOfOutputColors = 5
)

// TODO: document
// TODO: maybe provide a description string for showing inside the help
type RandomGenerator struct {
	Seed int64
}

func (pg *RandomGenerator) Description() string {
	return "Returns colors at a random offset from the first color"
}

func (pg *RandomGenerator) Generate(colors []*color.RGBA) []*color.RGBA {
	rand.Seed(pg.Seed)

	c := colors[0]
	value := float32(c.R+c.G+c.B) / 3

	out := make([]*color.RGBA, numberOfOutputColors)
	for i := 0; i < numberOfOutputColors; i++ {
		offset := rand.Float32()
		newValue := value + 2*rand.Float32()*offset - offset
		ratio := newValue / value

		red := float32(c.R) * ratio
		green := float32(c.G) * ratio
		blue := float32(c.B) * ratio
		newColor := &color.RGBA{
			R: uint8(255 * red),
			G: uint8(255 * green),
			B: uint8(255 * blue),
			A: 1}

		out[i] = newColor
	}
	return out
}
