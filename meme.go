package main

import (
	"flag"
	"os"
	"path"

	"github.com/fogleman/gg"
	"github.com/pelletier/go-toml"
)

func main() {
	// Show help
	if len(os.Args) < 2 || len(os.Args) > 3 || os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h" {
		println("Usage:")
		println("  meme [input] [output]")
		println("    input = an input .toml file for a meme")
		println("    output = a .png url to save the meme to")
		println("Example:")
		println("  meme ./foo.toml ./bar.png")
		return
	}

	input := os.Args[1]
	output := os.Args[2]
	flag.Parse()

	// Validate input
	if info, err := os.Stat(input); os.IsNotExist(err) || info.IsDir() {
		println("Please input a valid meme")
		return
	}

	// Validate output
	if _, err := os.Stat(path.Dir(output)); os.IsNotExist(err) {
		println("Please input a valid output file")
		return
	}

	// Load meme
	meme, err := toml.LoadFile(input)
	if err != nil {
		println("Error: Couldnt locate input meme")
		panic(err.Error())
	}

	// Load image
	image, err := gg.LoadImage(path.Join(path.Dir(input), meme.Get("image").(string)))
	if err != nil {
		println("Error: Couldnt load input meme")
		panic(err.Error())
	}

	ctx := gg.NewContextForImage(image)

	// Load font
	if err = ctx.LoadFontFace(
		meme.Get("font").(string),
		float64(meme.Get("fontSize").(int64)),
	); err != nil {
		println("Error: Couldnt load font")
		panic(err.Error())
	}

	// Draw text
	ctx.SetRGB(0, 0, 0)
	ctx.DrawStringAnchored("test", 310, 80, 0, 1)

	// Save image
	if err = ctx.SavePNG(output); err != nil {
		println("Error: Couldnt save output meme")
		panic(err.Error())
	}

	println("Meme saved to " + output)
	return
}
