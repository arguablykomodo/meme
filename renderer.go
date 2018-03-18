package main

import (
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
)

func render(input, output string, meme Meme, template Template) {
	templateDir := filepath.Dir(relative(input, meme.Template))
	memeDir := filepath.Dir(input)

	image, err := gg.LoadImage(filepath.Join(templateDir, template.Image))
	handleErr(err)

	ctx := gg.NewContextForImage(image)

	if filepath.IsAbs(template.Font) {
		err = ctx.LoadFontFace(template.Font, float64(template.FontSize))
		handleErr(err)
	} else {
		err = ctx.LoadFontFace(filepath.Join(templateDir, template.Font), float64(template.FontSize))
		handleErr(err)
	}

	ctx.SetRGB(0, 0, 0)
	for _, field := range template.Fields {
		if _, exists := meme.Fields[field.Name]; exists {
			text := meme.Fields[field.Name]
			if strings.HasPrefix(text, "url:") {
				imagePath := filepath.Join(memeDir, strings.TrimPrefix(text, "url:"))
				img, err := gg.LoadImage(imagePath)
				handleErr(err)

				scaleX := float64(field.W) / float64(img.Bounds().Size().X)
				scaleY := float64(field.H) / float64(img.Bounds().Size().Y)
				ctx.ScaleAbout(scaleX, scaleY, float64(field.X), float64(field.Y))
				ctx.DrawImage(img, field.X, field.Y)
				ctx.ScaleAbout(1/scaleX, 1/scaleY, float64(field.X), float64(field.Y))
			} else {
				var align gg.Align
				switch field.Align {
				case 0:
					align = gg.AlignLeft
				case 1:
					align = gg.AlignCenter
				case 2:
					align = gg.AlignRight
				}
				ctx.DrawStringWrapped(
					text,
					float64(field.X),
					float64(field.Y),
					0, 0,
					float64(field.W),
					1.25,
					align,
				)
			}
		}
	}

	err = ctx.SavePNG(output)
	handleErr(err)
}
