// TODO: document
package main

import (
	"github.com/snleu/poethe/poethe"

	"flag"
	"fmt"
	"os"
	"time"
)

const (
	defaultInputRGB = "#ef476f"

	defaultGeneratorID = "random"
)

func main() {
	inputColorRGB := flag.String("in", defaultInputRGB, "CSV list of RGB values")
	generatorID := flag.String("gen", defaultGeneratorID, "ID for the color generator")
	flag.Parse()

	if inputColorRGB == nil {
		fmt.Printf("Please provide a CSV of RGB values\n")
		os.Exit(1)
	}

	if generatorID == nil {
		fmt.Printf("Please provide an ID for the color generator\n")
		os.Exit(1)
	}

	// TODO: check if the inputColorRGB is not set to nil, get rid of panic in the following function
	colors, err := poethe.ParseColors(*inputColorRGB)
	if err != nil {
		fmt.Printf("Could not parse color codes: %v\n", err)
		os.Exit(1)
	}

	config := &appConfiguration{
		InputColorRGB: *inputColorRGB,
	}
	if _, err := validConfig(config); err != nil {
		fmt.Printf("Invalid input: %v\n", err)
		os.Exit(1)
	}

	generators := buildGenerators()
	gen, ok := generators[*generatorID]
	if !ok {
		fmt.Printf("No generator found for key: %s\n", *generatorID)
		os.Exit(1)
	}

	palette := gen.Generate(colors)

	// Generate the colors
	formattedOutput := poethe.FormatColors(palette)
	fmt.Printf("%s", formattedOutput)
}

type appConfiguration struct {
	InputColorRGB string
}

func validConfig(c *appConfiguration) (bool, error) {
	if c == nil {
		panic(fmt.Errorf("No configuration provided"))
	}

	if len(c.InputColorRGB) == 0 {
		return false, fmt.Errorf("No input color RGB provided")
	}

	return true, nil
}

func buildGenerators() map[string]poethe.Generator {
	generators := make(map[string]poethe.Generator)
	rg := &poethe.RandomGenerator{
		Seed: time.Now().UnixNano(),
	}
	generators["random"] = rg

	return generators
}
