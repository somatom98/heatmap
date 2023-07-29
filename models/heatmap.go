package models

import (
	"github.com/somatom98/heatmap/constants"
)

type Heatmap[T any] struct {
	Width   int
	Height  int
	Squares []HeatSquareGroup[T]
}

type HeatSquareGroup[T any] struct {
	Direction constants.Direction
	Squares   []HeatSquare[T]
}
