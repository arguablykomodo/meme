package main

import (
	"flag"
	"os"
	"path"
)

func main() {
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

	if info, err := os.Stat(input); os.IsNotExist(err) || info.IsDir() {
		println("Please input a valid meme")
		return
	}

	if _, err := os.Stat(path.Dir(output)); os.IsNotExist(err) {
		println("Please input a valid output file")
		return
	}

	println(input, output)
}
