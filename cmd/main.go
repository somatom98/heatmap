package main

import (
	"fmt"
	"log"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/somatom98/heatmap"
	"github.com/somatom98/heatmap/models"
	"github.com/somatom98/heatmap/repositories"
	"github.com/somatom98/heatmap/services"
)

var (
	currencyRepository     heatmap.HeatsquareRepository[models.Currency]
	currencyHeatmapService heatmap.HeatmapService[models.Currency]
)

func main() {
	currencyRepository = repositories.NewCurrencyRepository()
	currencyHeatmapService = services.NewHeatmapService[models.Currency](currencyRepository)

	http.Handle("/heatmap", http.HandlerFunc(heatmapHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func heatmapHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	heatmap := currencyHeatmapService.Create()

	s := svg.New(w)
	s.Start(heatmap.Width, heatmap.Height)
	s.Rect(0, 0, heatmap.Width, heatmap.Height, "stroke:black")
	for _, square := range heatmap.Squares {
		currencyRect(s, square.X, square.Y, square.Width, square.Height, square.Info)
	}
	s.End()
}

func currencyRect(s *svg.SVG, x, y, w, h int, currency models.Currency) {
	color := "green"
	if currency.Last24Variation < 0 {
		color = "red"
	}
	s.Rect(x, y, w, h, "fill:"+color+";stroke:black")
	s.Text(x+10, y+20, currency.Name, "font-size:14pt;fill:white")
	s.Text(x+10, y+40, fmt.Sprintf("%.2f", currency.Price), "font-size:14pt;fill:white")
	s.Text(x+10, y+60, fmt.Sprintf("%.2f", currency.Last24Variation), "font-size:14pt;fill:white")
}
