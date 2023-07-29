package services

import (
	"github.com/somatom98/heatmap"
	"github.com/somatom98/heatmap/models"
)

const (
	Width  = 1000
	Height = 1000
)

type HeatmapService[T heatmap.Value] struct {
	repository heatmap.HeatsquareRepository[T]
	heatmap    models.Heatmap[T]
}

func NewHeatmapService[T heatmap.Value](repository heatmap.HeatsquareRepository[T]) *HeatmapService[T] {
	return &HeatmapService[T]{
		repository: repository,
	}
}

func (s *HeatmapService[T]) Create() models.Heatmap[T] {
	s.heatmap = models.Heatmap[T]{
		Width:  Width,
		Height: Height,
	}

	values := s.repository.GetAll()
	for i, value := range values {
		s.heatmap.Squares = append(s.heatmap.Squares, models.HeatSquare[T]{
			X:      10,
			Y:      10 + i*110,
			Width:  200,
			Height: 100,
			Info:   value,
			Color:  value.Color(),
		})
	}

	return s.heatmap
}
