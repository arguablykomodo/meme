package main

type meme struct {
	Template string
	Fields   map[string]string
}

type template struct {
	Image    string
	HAlign   int
	VAlign   int
	Font     string
	FontSize float64
	Color    []float64
	Rotation float64
	Fields   []field
}

type field struct {
	Name     string
	HAlign   int
	VAlign   int
	FontSize float64
	Color    []float64
	Rotation float64
	X        float64
	Y        float64
	W        float64
	H        float64
}
