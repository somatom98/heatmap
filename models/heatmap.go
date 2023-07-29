package models

type Heatmap[T any] struct {
	Width   int
	Height  int
	Squares []HeatSquare[T]
}
