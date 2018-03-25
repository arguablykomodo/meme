package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Show help
	if len(os.Args) < 2 || len(os.Args) > 3 || os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println(
			`Usage:
  meme [input] [output]
    input = an input .toml file for a meme
    output = a .png url to save the meme to
Example:
  meme foo.toml bar.png`,
		)
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	// Validate input
	if info, err := os.Stat(input); os.IsNotExist(err) || info.IsDir() {
		fmt.Println("Please input a valid meme")
		return
	}

	// Validate output
	if _, err := os.Stat(filepath.Dir(output)); os.IsNotExist(err) {
		fmt.Println("Please input a valid output file")
		return
	}

	render(input, output)

	fmt.Println("Meme saved to " + output)
	return
}
