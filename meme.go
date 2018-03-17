package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/fogleman/gg"
)

func relative(path1, path2 string) string {
	return path.Join(path.Dir(path1), path2)
}

func main() {
	// Show help
	if len(os.Args) < 2 || len(os.Args) > 3 || os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println("Usage:")
		fmt.Println("  meme [input] [output]")
		fmt.Println("    input = an input .toml file for a meme")
		fmt.Println("    output = a .png url to save the meme to")
		fmt.Println("Example:")
		fmt.Println("  meme ./foo.toml ./bar.png")
		return
	}

	input := os.Args[1]
	output := os.Args[2]
	flag.Parse()

	// Validate input
	if info, err := os.Stat(input); os.IsNotExist(err) || info.IsDir() {
		fmt.Println("Please input a valid meme")
		return
	}

	// Validate output
	if _, err := os.Stat(path.Dir(output)); os.IsNotExist(err) {
		fmt.Println("Please input a valid output file")
		return
	}

	// Load meme
	var meme Meme
	if _, err := toml.DecodeFile(input, &meme); err != nil {
		fmt.Println("Error: Couldnt load meme file")
		panic(err.Error())
	}

	// Load template
	var template Template
	if _, err := toml.DecodeFile(relative(input, meme.Template), &template); err != nil {
		fmt.Println("Error: Couldnt load template file")
		panic(err.Error())
	}

	// Load image
	image, err := gg.LoadImage(relative(relative(input, meme.Template), template.Image))
	if err != nil {
		fmt.Println("Error: Couldnt load meme image")
		panic(err.Error())
	}

	ctx := gg.NewContextForImage(image)

	// Load font
	if err = ctx.LoadFontFace(
		template.Font,
		float64(template.FontSize),
	); err != nil {
		fmt.Println("Error: Couldnt load font")
		panic(err.Error())
	}

	// Draw text
	ctx.SetRGB(0, 0, 0)
	for field, text := range meme.Fields {
		templateField := template.Fields[field]
		for _, box := range templateField.Coords {
			ctx.DrawStringWrapped(
				text,
				float64(box.X),
				float64(box.Y),
				0, 0,
				float64(box.W),
				float64(template.FontSize)/10,
				gg.AlignLeft,
			)
		}
	}

	// Save image
	if err = ctx.SavePNG(output); err != nil {
		fmt.Println("Error: Couldnt save output meme")
		panic(err.Error())
	}

	fmt.Println("Meme saved to " + output)
	return
}
