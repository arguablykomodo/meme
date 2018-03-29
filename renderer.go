package main

import (
	"image"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/fogleman/gg"
)

func drawImage(img image.Image, ctx *gg.Context, x, y, w, h float64) {
	// Calculate the scaling factor
	scaleX := w / float64(img.Bounds().Size().X)
	scaleY := h / float64(img.Bounds().Size().Y)
	// Scale and draw
	ctx.ScaleAbout(scaleX, scaleY, x, y)
	ctx.DrawImage(img, int(x), int(y))
	// Reverse the scaling
	ctx.ScaleAbout(1/scaleX, 1/scaleY, x, y)
}

func render(input string, i int) image.Image {
	if i > 10 {
		return gg.NewContext(1, 1).Image()
	}

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
				path := resolvePath(strings.TrimPrefix(text, "url:"), memeDir)
				if filepath.Ext(path) == ".toml" {
					img := render(path, i+1)
					drawImage(img, ctx, float64(field.X), float64(field.Y), float64(field.W), float64(field.H))
				} else {
					img, err := gg.LoadImage(path)
					handleErr(err)
					drawImage(img, ctx, float64(field.X), float64(field.Y), float64(field.W), float64(field.H))
				}
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
	return ctx.Image()
}
