package main

import (
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/fogleman/gg"
)

func render(input, output string) {
	// Load meme
	var meme Meme
	_, err := toml.DecodeFile(input, &meme)
	handleErr(err)

	// Util variables for directories
	memeDir := filepath.Dir(input)
	templateDir := filepath.Dir(resolvePath(meme.Template, memeDir))

	// Load template
	var template Template
	_, err = toml.DecodeFile(resolvePath(meme.Template, memeDir), &template)
	handleErr(err)

	// Get source image
	image, err := gg.LoadImage(resolvePath(template.Image, templateDir))
	handleErr(err)

	// Create context
	ctx := gg.NewContextForImage(image)

	// Load font
	err = ctx.LoadFontFace(resolvePath(template.Font, templateDir), float64(template.FontSize))
	handleErr(err)

	// Set color
	ctx.SetRGB(template.Color[0], template.Color[1], template.Color[2])

	// For each field in the template
	for _, field := range template.Fields {
		// If the meme has that field
		if text, exists := meme.Fields[field.Name]; exists {
			// If we need to load an image
			if strings.HasPrefix(text, "url:") {
				// Load the image
				img, err := gg.LoadImage(resolvePath(strings.TrimPrefix(text, "url:"), memeDir))
				handleErr(err)

				// Calculate the scaling factor
				scaleX := float64(field.W) / float64(img.Bounds().Size().X)
				scaleY := float64(field.H) / float64(img.Bounds().Size().Y)
				// Scale and draw
				ctx.ScaleAbout(scaleX, scaleY, float64(field.X), float64(field.Y))
				ctx.DrawImage(img, field.X, field.Y)
				// Reverse the scaling
				ctx.ScaleAbout(1/scaleX, 1/scaleY, float64(field.X), float64(field.Y))
			} else { // If we need to render a string
				// Then draw the string
				ctx.DrawStringWrapped(
					text,
					float64(field.X),
					float64(field.Y),
					0, 0,
					float64(field.W),
					1.25,
					gg.Align(field.Align),
				)
			}
		}
	}

	// Save the image
	err = ctx.SavePNG(output)
	handleErr(err)
}
