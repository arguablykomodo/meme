package main

import (
	"fmt"
	"path/filepath"

	"github.com/fogleman/gg"
)

func render(input, output string, meme Meme, template Template) {
	image, err := gg.LoadImage(relative(relative(input, meme.Template), template.Image))
	handleErr(err)

	ctx := gg.NewContextForImage(image)

	fmt.Println(template.Font, filepath.IsAbs(template.Font))
	if filepath.IsAbs(template.Font) {
		err = ctx.LoadFontFace(template.Font, float64(template.FontSize))
		handleErr(err)
	} else {
		err = ctx.LoadFontFace(relative(relative(input, meme.Template), template.Font), float64(template.FontSize))
		handleErr(err)
	}

	ctx.SetRGB(0, 0, 0)
	for field, text := range meme.Fields {
		field := template.Fields[field]
		ctx.DrawStringWrapped(
			text,
			float64(field.X),
			float64(field.Y),
			0, 0,
			float64(field.W),
			float64(template.FontSize)/10,
			gg.AlignLeft,
		)
	}

	err = ctx.SavePNG(output)
	handleErr(err)
}
