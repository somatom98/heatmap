package services

import (
	"fmt"
	"io"
	"math"

	svg "github.com/ajstarks/svgo"
	"github.com/somatom98/heatmap"
	"github.com/somatom98/heatmap/models"
)

type HeatmapDrawerService[T heatmap.Value] struct {
	heatmapService heatmap.HeatmapService[T]
	canvas         *svg.SVG
}

func NewHeatmapDrawerService[T heatmap.Value](heatmapService heatmap.HeatmapService[T]) *HeatmapDrawerService[T] {
	return &HeatmapDrawerService[T]{
		heatmapService: heatmapService,
	}
}

func (s *HeatmapDrawerService[T]) Draw(w io.Writer) {
	heatmap := s.heatmapService.Create()

	s.canvas = svg.New(w)
	s.canvas.Start(heatmap.Width, heatmap.Height)

	s.canvas.Rect(0, 0, heatmap.Width, heatmap.Height, "fill:white")
	for _, row := range heatmap.Squares {
		for _, square := range row.Squares {
			s.currencyRect(square)
		}
	}
	s.canvas.End()
}

func (s *HeatmapDrawerService[T]) currencyRect(square models.HeatSquare[T]) {
	middleX := square.X + (square.Width / 2)
	middleY := square.Y + (square.Height / 2)

	fontSize := int(math.Min(float64(square.Width), float64(square.Height)) / 5)
	subFontSize := fontSize / 2

	s.canvas.Rect(square.X, square.Y, square.Width, square.Height, "fill:"+square.Color+";stroke:black")
	s.canvas.Text(middleX, middleY, square.Info.DisplayName(), s.textStyle(fontSize))
	s.canvas.Text(middleX, middleY+2+subFontSize, square.Info.DisplayValue(), s.textStyle(subFontSize))
	s.canvas.Text(middleX, middleY+2+subFontSize*2, square.Info.DisplaySubValue(), s.textStyle(subFontSize))
}

func (s *HeatmapDrawerService[T]) textStyle(fontSize int) string {
	return fmt.Sprintf("font-size:%dpt;font-family:Arial;fill:white;text-anchor:middle", fontSize)
}
