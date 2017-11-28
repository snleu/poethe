// TODO: document
package poethe

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"
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

func (pg *RandomGenerator) Generate(colors []*color.RGBA) ([]*color.RGBA, error) {

	if len(colors) == 0 {
		return nil, fmt.Errorf("RandomGenerator needs at least 1 color")
	}
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
	return out, nil
}

type TriadMixGenerator struct {
	GreyControl float32
}

func (tg *TriadMixGenerator) Description() string {
	return "Takes three colours, and mixes them randomly to create a palette."
}

func (tg *TriadMixGenerator) Generate(colors []*color.RGBA) ([]*color.RGBA, error) {

	if len(colors) != 3 {
		return nil, fmt.Errorf("TriadMixGenerator needs 3 colors")
	}

	rand.Seed(time.Now().UnixNano())
	out := make([]*color.RGBA, numberOfOutputColors)
	for i := 0; i < numberOfOutputColors; i++ {
		out[i] = tg.generateTriadMix(colors)
	}

	return out, nil
}

func (tg *TriadMixGenerator) generateTriadMix(colors []*color.RGBA) *color.RGBA {
	randomIndex := rand.Int() % 3

	var mixRatio0 float32
	var mixRatio1 float32
	var mixRatio2 float32

	if randomIndex == 0 {
		mixRatio0 = rand.Float32() * tg.GreyControl
	} else {
		mixRatio0 = rand.Float32()
	}
	if randomIndex == 1 {
		mixRatio1 = rand.Float32() * tg.GreyControl
	} else {
		mixRatio1 = rand.Float32()
	}
	if randomIndex == 2 {
		mixRatio2 = rand.Float32() * tg.GreyControl
	} else {
		mixRatio2 = rand.Float32()
	}
	sum := mixRatio0 + mixRatio1 + mixRatio2

	mixRatio0 = mixRatio0 / sum
	mixRatio1 = mixRatio1 / sum
	mixRatio2 = mixRatio2 / sum

	newColor := &color.RGBA{
		R: uint8(mixRatio0*float32(colors[0].R) + mixRatio1*float32(colors[1].R) + mixRatio2*float32(colors[2].R)),
		G: uint8(mixRatio0*float32(colors[0].G) + mixRatio1*float32(colors[1].G) + mixRatio2*float32(colors[2].G)),
		B: uint8(mixRatio0*float32(colors[0].B) + mixRatio1*float32(colors[1].B) + mixRatio2*float32(colors[2].B)),
		A: 255,
	}
	return newColor
}
