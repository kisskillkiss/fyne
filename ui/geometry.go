package ui

type Size struct {
	Width  int
	Height int
}

func NewSize(w int, h int) Size {
	return Size{w, h}
}

type Position struct {
	X int
	Y int
}

func (p1 Position) Add(p2 Position) Position {
	return Position{p1.X + p2.X, p1.Y + p2.Y}
}

func NewPos(x int, y int) Position {
	return Position{x, y}
}
