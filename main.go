package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Show help if commands are invalid
	if len(os.Args) != 2 {
		fmt.Println(
			`Usage:
  meme [file]
    the program will render the meme at [file]
  meme [dir]
    the program will render all the memes in [dir]`,
		)
		return
	}

	input := os.Args[1]

	// Validate input
	info, err := os.Stat(input)
	if os.IsNotExist(err) {
		fmt.Println("there is no meme/directory at the input location")
		return
	}
	handleErr(err)

	// If input is a dir
	if info.IsDir() {
		// Render all files in that dir
		files, err := ioutil.ReadDir(input)
		handleErr(err)

		for _, file := range files {
			if !file.IsDir() {
				inFile := filepath.Join(input, file.Name())
				output := strings.TrimSuffix(inFile, filepath.Ext(file.Name())) + ".png"
				render(inFile, output)
				fmt.Println(inFile + " saved to " + output)
			}
		}

	} else { // If it is a file
		// Then just render it
		output := strings.TrimSuffix(input, filepath.Ext(input)) + ".png"
		render(input, output)
		fmt.Println(input + " saved to " + output)
	}

	return
}
