package main

type meme struct {
	Template string
	Fields   map[string]string
}

type template struct {
	Image    string
	Font     string
	FontSize float64
	Color    []float64
	Fields   []field
}

type field struct {
	Name  string
	Align int
	X     float64
	Y     float64
	W     float64
	H     float64
}
