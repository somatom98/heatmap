package services

import (
	"math"
	"sort"

	"github.com/somatom98/heatmap"
	"github.com/somatom98/heatmap/constants"
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
			lastSquare = s.newSquare(value, ratio, lastSquare, group)
			group.Squares = append(group.Squares, lastSquare)
			continue
		}

		spaceLeft := s.spaceLeft(lastSquare, group.Direction)
		if spaceLeft < value.Area()*ratio {
			group = s.fillGaps(group, ratio)
			s.heatmap.Squares = append(s.heatmap.Squares, group)
			lastSquare = s.newSquare(value, ratio, lastSquare, group)
			group = s.newGroup(group)
			group.Squares = append(group.Squares, lastSquare)
			continue
		}

		lastSquare = s.newGroupSquare(value, ratio, lastSquare, group.Direction)
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

	startX, startY := s.startAt(group)

	newHeight := totalVolume / float64(Width-startX)
	newWidth := totalVolume / float64(Height-startY)

	newGroup := models.HeatSquareGroup[T]{
		Direction: group.Direction,
		Squares:   make([]models.HeatSquare[T], 0),
	}

	lastX, lastY := startX, startY
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

func (s *HeatmapService[T]) startAt(group models.HeatSquareGroup[T]) (int, int) {
	if len(group.Squares) == 0 {
		return 0, 0
	}
	if group.Direction == constants.Horizontal {
		return group.Squares[0].X, group.Squares[0].Y + group.Squares[0].Height
	}
	return group.Squares[0].X + group.Squares[0].Width, group.Squares[0].Y
}

func (s *HeatmapService[T]) newGroup(group models.HeatSquareGroup[T]) models.HeatSquareGroup[T] {
	startX, startY := s.startAt(group)

	newGroup := models.HeatSquareGroup[T]{
		Squares: make([]models.HeatSquare[T], 0),
	}

	// choose shorter side
	if Width-startX < Height-startY {
		newGroup.Direction = constants.Horizontal
		return newGroup
	}
	newGroup.Direction = constants.Vertical
	return newGroup
}

func (s *HeatmapService[T]) newGroupSquare(value T, ratio float64, lastSquare models.HeatSquare[T], direction constants.Direction) models.HeatSquare[T] {
	if direction == constants.Horizontal {
		return models.HeatSquare[T]{
			Color:  value.Color(),
			X:      lastSquare.X + lastSquare.Width,
			Y:      lastSquare.Y,
			Width:  int(value.Area() * ratio / float64(lastSquare.Height)),
			Height: int(lastSquare.Height),
			Info:   value,
		}
	}
	return models.HeatSquare[T]{
		Color:  value.Color(),
		X:      lastSquare.X,
		Y:      lastSquare.Y + lastSquare.Height,
		Width:  int(lastSquare.Width),
		Height: int(value.Area() * ratio / float64(lastSquare.Width)),
		Info:   value,
	}
}

func (s *HeatmapService[T]) newSquare(value T, ratio float64, lastSquare models.HeatSquare[T], group models.HeatSquareGroup[T]) models.HeatSquare[T] {
	startX, startY := s.startAt(group)
	return models.HeatSquare[T]{
		Color:  value.Color(),
		X:      startX,
		Y:      startY,
		Width:  int(math.Sqrt(value.Area() * ratio * Width / Height)),
		Height: int(math.Sqrt(value.Area() * ratio * Height / Width)),
		Info:   value,
	}
}
