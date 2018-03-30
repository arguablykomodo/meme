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
	if i > 10 { // Limit recursion to 10 steps
		return gg.NewContext(1, 1).Image()
	}

	// Load everything
	var meme meme
	_, err := toml.DecodeFile(input, &meme)
	handleErr(err)

	memeDir := filepath.Dir(input)
	templateDir := filepath.Dir(resolvePath(meme.Template, memeDir))

	var template template
	_, err = toml.DecodeFile(resolvePath(meme.Template, memeDir), &template)
	handleErr(err)

	image, err := gg.LoadImage(resolvePath(template.Image, templateDir))
	handleErr(err)

	ctx := gg.NewContextForImage(image)

	err = ctx.LoadFontFace(resolvePath(template.Font, templateDir), template.FontSize)
	handleErr(err)

	ctx.SetRGB(template.Color[0], template.Color[1], template.Color[2])

	for _, field := range template.Fields {
		if text, exists := meme.Fields[field.Name]; exists { // For each field in the meme

			switch {
			case strings.HasPrefix(text, "text:"): // Just draw the text if its a text field
				ctx.DrawStringWrapped(strings.TrimPrefix(text, "text:"), field.X, field.Y, 0, 0, field.W, 1.25, gg.Align(field.Align))

			case strings.HasPrefix(text, "url:"): // If it is an url then draw the image/meme at that location
				path := resolvePath(strings.TrimPrefix(text, "url:"), memeDir)
				switch filepath.Ext(path) {
				case ".toml":
					drawImage(render(path, i+1), ctx, field.X, field.Y, field.W, field.H)
				default:
					img, err := gg.LoadImage(path)
					handleErr(err)
					drawImage(img, ctx, field.X, field.Y, field.W, field.H)
				}
			}
		}
	}

	// Return the image
	return ctx.Image()
}
