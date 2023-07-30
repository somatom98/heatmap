package services

import (
	"math"
	"sort"

	"github.com/somatom98/heatmap"
	"github.com/somatom98/heatmap/constants"
	"github.com/somatom98/heatmap/models"
)

const (
	Width  = 800
	Height = 800
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
	ratio := s.areaRatio(values)
	sort.Slice(values, func(i, j int) bool {
		return values[i].Area() > values[j].Area()
	})

	s.groupSquares(values, ratio)

	return s.heatmap
}

func (s *HeatmapService[T]) areaRatio(values []T) float64 {
	totalArea := 0.0
	for _, value := range values {
		totalArea += value.Area()
	}

	heatmapArea := float64(Width * Height)
	return heatmapArea / totalArea
}

func (s *HeatmapService[T]) groupSquares(values []T, ratio float64) {
	group := models.HeatSquareGroup[T]{
		Direction: constants.Horizontal,
		Squares:   make([]models.HeatSquare[T], 0),
	}

	lastSquare := models.HeatSquare[T]{}
	for _, value := range values {
		if len(group.Squares) == 0 {
			lastSquare = models.HeatSquare[T]{
				Color:  value.Color(),
				X:      0,
				Y:      0,
				Width:  int(math.Sqrt(value.Area() * ratio * Width / Height)),
				Height: int(math.Sqrt(value.Area() * ratio * Height / Width)),
				Info:   value,
			}
			group.Squares = append(group.Squares, lastSquare)
			continue
		}

		spaceLeft := s.spaceLeft(lastSquare, group.Direction)
		if spaceLeft < value.Area()*ratio {
			s.heatmap.Squares = append(s.heatmap.Squares, s.fillGaps(group, ratio))
			group = models.HeatSquareGroup[T]{
				Direction: constants.Horizontal,
				Squares:   make([]models.HeatSquare[T], 0),
			}
			lastSquare = models.HeatSquare[T]{
				Color:  value.Color(),
				X:      0,
				Y:      lastSquare.Y + s.heatmap.Squares[len(s.heatmap.Squares)-1].Squares[len(s.heatmap.Squares[len(s.heatmap.Squares)-1].Squares)-1].Height,
				Width:  int(math.Sqrt(value.Area() * ratio * Width / Height)),
				Height: int(math.Sqrt(value.Area() * ratio * Height / Width)),
				Info:   value,
			}
			group.Squares = append(group.Squares, lastSquare)
			continue
		}

		lastSquare = models.HeatSquare[T]{
			Color:  value.Color(),
			X:      lastSquare.X + lastSquare.Width,
			Y:      lastSquare.Y,
			Width:  int(value.Area() * ratio / float64(lastSquare.Height)),
			Height: int(lastSquare.Height),
			Info:   value,
		}
		group.Squares = append(group.Squares, lastSquare)
	}

	s.heatmap.Squares = append(s.heatmap.Squares, s.fillGaps(group, ratio))
}

func (s *HeatmapService[T]) spaceLeft(square models.HeatSquare[T], direction constants.Direction) float64 {
	if direction == constants.Horizontal {
		return float64((Width - square.X - square.Width) * square.Height)
	}
	return float64((Height - square.Y - square.Height) * square.Width)
}

func (s *HeatmapService[T]) fillGaps(group models.HeatSquareGroup[T], ratio float64) models.HeatSquareGroup[T] {
	totalVolume := 0.0
	for _, square := range group.Squares {
		totalVolume += square.Info.Area() * ratio
	}

	newHeight := totalVolume / float64(Width)
	newWidth := totalVolume / float64(Height)

	newGroup := models.HeatSquareGroup[T]{
		Direction: group.Direction,
		Squares:   make([]models.HeatSquare[T], 0),
	}

	lastX, lastY := 0, 0
	for _, square := range group.Squares {
		if group.Direction == constants.Horizontal {
			square.X = lastX
			square.Height = int(newHeight)
			square.Width = int(square.Info.Area() * ratio / newHeight)
			lastX = square.X + square.Width
		} else {
			square.Y = lastY
			square.Width = int(newWidth)
			square.Height = int(square.Info.Area() * ratio / newWidth)
			lastY = square.Y + square.Height
		}
		newGroup.Squares = append(newGroup.Squares, square)
	}

	return newGroup
}
