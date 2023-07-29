package heatmap

import (
	svg "github.com/ajstarks/svgo"
	"github.com/somatom98/heatmap/models"
)

type HeatsquareRepository[T any] interface {
	GetAll() []models.HeatSquare[T]
}

type HeatmapService[T any] interface {
	Create() models.Heatmap[T]
}

type HeatmapDrawer[T any] interface {
	Draw(models.Heatmap[T]) *svg.SVG
}
