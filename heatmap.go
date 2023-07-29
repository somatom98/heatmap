package heatmap

import (
	"io"

	"github.com/somatom98/heatmap/models"
)

type Value interface {
	Color() string
	Area() float64
	DisplayName() string
	DisplayValue() string
	DisplaySubValue() string
}

type HeatsquareRepository[T Value] interface {
	GetAll() []T
}

type HeatmapService[T Value] interface {
	Create() models.Heatmap[T]
}

type HeatmapDrawerService[T Value] interface {
	Draw(io.Writer)
}
