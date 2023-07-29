package heatmap

import (
	svg "github.com/ajstarks/svgo"
	"github.com/somatom98/heatmap/models"
)

type Value interface {
	Color() string
}

type HeatsquareRepository[T Value] interface {
	GetAll() []T
}

type HeatmapService[T Value] interface {
	Create() models.Heatmap[T]
}

type HeatmapDrawer[T Value] interface {
	Draw(models.Heatmap[T]) *svg.SVG
}
