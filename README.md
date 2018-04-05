# The meme-making CLI you never asked for!
Finally a command line interface for automating shitty meme templates

Licensed under CC0 1.0 Universal, so that anyone everywhere can use and abuse this monster that i have created

## Why did you make this?
I did this due to a combination of several different things

1. Boredom
2. Free time
3. Wanting to learn Go
4. ~~Crippling depression~~ Love

## How do you use this?

### Glossary
| Name     | Meaning                                      |
|----------|----------------------------------------------|
| `meme`   | The program                                  |
| Meme     | A .toml file that represents a meme          |
| Template | A .toml file that represents a meme template |

### CLI
`meme` has the following commands

- `meme [file]`
  - `[file]` would be a Meme that you want to render into an image
  - For example `meme thing.toml` would render the Meme in `thing.toml` into an image at `thing.png`
- `meme [directory]`
  - `[directory]` would be a directory containing several Meme files: These files will be all rendered into an image
  - For example `meme .` would render all Memes in the current directory

### Schema

#### Template
A Template is a [TOML](https://github.com/toml-lang/toml) encoded file, it must contain these properties:

| Name     | Meaning                                                                                                                |
|---------:|:-----------------------------------------------------------------------------------------------------------------------|
| Image    | A path to an image file, this will be the base template of the Meme                                                    |
| HAlign   | The horizontal alignment for the text in the Meme can be 1 for left aligned, 2 for center align, and 3 for right align |
| VAlign   | Same as HAlign, but for vertical alignment                                                                             |
| Font     | A path to a font file, this will be the font that the text in the Meme will use                                        |
| FontSize | Pretty self-explanatory                                                                                                |
| Color    | Defines the color that the text will use via RGB values from 0 to 1. For example [1.0, 1.0, 1.0] for white             |
|	Rotation | Defines the rotation of the text in degrees                                                                            |

Besides all of these properties, a Template has the `Fields` property, wich is an array of structs that have the following properties:

| Name | Meaning                                                                                                                                    |
|-----:|:-------------------------------------------------------------------------------------------------------------------------------------------|
| Name | An identifier for the field, this will be used in the Meme file for putting text or an image in it, multiple fields can have the same name |
| X    | The X position of the field                                                                                                                |
| Y    | The Y position of the field                                                                                                                |
| W    | The width of the field                                                                                                                     |
| H    | The height of the field                                                                                                                    |

You can also use the same HAlign/VAlign/Font/FontSize/Color options from the Template in each field, and those settings will be applied individually to that field

For example, this is a simple template file for the good old Drake meme

```toml
Image="drake.png"
Font="C:/Windows/Fonts/arial.ttf"
FontSize=50.0
Color=[0.0, 0.0, 0.0]
HAlign=2
VAlign=2

[[Fields]]
Name="Bad"
X=674.0
Y=0.0
W=670.0
H=670.0

[[Fields]]
Name="Good"
X=674.0
Y=674.0
W=670.0
H=670.0
```

#### Meme
A Meme is also a TOML encoded file, and it is the one that "implements" the Template that you defined earlier.

It has a `Template` property, wich is a path to the Template that you want to implement

The other property is the `Fields` property, for each field in the Template, you have a property in `Fields` with that field's name in the Meme. For example, if your Template has a field named "Foo" and you want it to have the text "Bar", you would add `Foo="text:Bar"` to your Meme.

As you probably noticed, the text has the `text:` prefix, this let's the program know you are writing text. If you wanted to, for example, put an image called `bar.png` into that field, you would write `Foo="url:bar.png"`

This is getting a little complicated, so let's see an example of a Meme file for the Drake template we saw earlier:

```toml
Template="Templates/drake.toml"

[Fields]
  Bad="text:Making a meme normally"
  Good="text:Creating a way too complicated program for making them automatically"
```
