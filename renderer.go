package main

import (
	"image"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/fogleman/gg"
)

func drawImage(img image.Image, ctx *gg.Context, x, y, w, h, r float64) {
	// Calculate the scaling factor
	scaleX := w / float64(img.Bounds().Size().X)
	scaleY := h / float64(img.Bounds().Size().Y)
	// Transform and draw
	ctx.Push()
	ctx.RotateAbout(r, x+w/2, y+h/2)
	ctx.ScaleAbout(scaleX, scaleY, x, y)
	ctx.DrawImage(img, int(x), int(y))
	ctx.Pop()
}

func drawText(text string, ctx *gg.Context, align gg.Align, x, y, w, h, r float64) {
	ctx.Push()
	ctx.RotateAbout(r, x+w/2, y+h/2)
	ctx.DrawStringWrapped(strings.TrimPrefix(text, "text:"), x, y, 0, 0, w, 1.25, align)
	ctx.Pop()
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

	for _, field := range template.Fields {
		if text, exists := meme.Fields[field.Name]; exists { // For each field in the meme

			align := gg.Align(template.Align - 1)
			if field.Align != 0 {
				align = gg.Align(field.Align - 1)
			}

			fontSize := template.FontSize
			if field.FontSize != 0 {
				fontSize = field.FontSize
			}
			err = ctx.LoadFontFace(resolvePath(template.Font, templateDir), fontSize)
			handleErr(err)

			color := template.Color
			if field.Color != nil {
				color = field.Color
			}
			ctx.SetRGB(color[0], color[1], color[2])

			rotation := template.Rotation
			if field.Rotation != 0 {
				rotation = field.Rotation
			}

			switch {
			case strings.HasPrefix(text, "text:"): // Just draw the text if its a text field
				drawText(text, ctx, align, field.X, field.Y, field.W, field.H, gg.Radians(rotation))

			case strings.HasPrefix(text, "url:"): // If it is an url then draw the image/meme at that location
				path := resolvePath(strings.TrimPrefix(text, "url:"), memeDir)
				switch filepath.Ext(path) {
				case ".toml":
					drawImage(render(path, i+1), ctx, field.X, field.Y, field.W, field.H, gg.Radians(rotation))
				default:
					img, err := gg.LoadImage(path)
					handleErr(err)
					drawImage(img, ctx, field.X, field.Y, field.W, field.H, field.Rotation)
				}
			}
		}
	}

	// Return the image
	return ctx.Image()
}
