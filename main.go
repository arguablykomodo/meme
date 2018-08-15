package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
)

func main() {
	// Show help if commands are invalid
	if len(os.Args) < 2 {
		fmt.Println(
			`Usage:
  meme [file, file2, ...fileN]
    the program will render the meme files specified
  meme [dir, dir2, ...dirN]
    the program will render all the memes in the specified folders`,
		)
		return
	}

	args := os.Args[1:]

	for _, input := range args {

		// Validate input
		info, err := os.Stat(input)
		if os.IsNotExist(err) {
			fmt.Println("there is no meme/directory at " + input)
			continue
		} else {
			handleErr(err)
		}

		// If input is a dir
		switch {
		case info.IsDir():
			// Render all files in that dir
			files, err := ioutil.ReadDir(input)
			handleErr(err)

			for _, file := range files {
				if !file.IsDir() && filepath.Ext(file.Name()) == ".toml" {
					inFile := filepath.Join(input, file.Name())
					output := strings.TrimSuffix(inFile, filepath.Ext(file.Name())) + ".png"
					image := render(inFile, 0)
					err = gg.SavePNG(output, image)
					handleErr(err)
					fmt.Println(inFile + " saved to " + output)
				}
			}

		case filepath.Ext(input) == ".toml": // If it is a file
			// Then just render it
			output := strings.TrimSuffix(input, filepath.Ext(input)) + ".png"
			image := render(input, 0)
			err = gg.SavePNG(output, image)
			handleErr(err)
			fmt.Println(input + " saved to " + output)

		default:
			fmt.Println("there is no meme/directory at " + input)
		}
	}
}
