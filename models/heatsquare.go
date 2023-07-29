package models

type HeatSquare[T any] struct {
	Info   T
	X      int
	Y      int
	Width  int
	Height int
	Color  string
}
