package main

// Field A field in the template
type Field struct {
	Name  string
	Align int
	X     int
	Y     int
	W     int
	H     int
}

// Template A meme template
type Template struct {
	Image    string
	Font     string
	FontSize int
	Fields   []Field
}

// Meme A meme
type Meme struct {
	Template string
	Fields   map[string]string
}
