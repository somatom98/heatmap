package main

import (
	"log"
	"net/http"

	"github.com/somatom98/heatmap"
	"github.com/somatom98/heatmap/models"
	"github.com/somatom98/heatmap/repositories"
	"github.com/somatom98/heatmap/services"
)

var (
	currencyRepository           heatmap.HeatsquareRepository[models.Currency]
	currencyHeatmapService       heatmap.HeatmapService[models.Currency]
	currencyHeatmapDrawerService heatmap.HeatmapDrawerService[models.Currency]
)

func main() {
	currencyRepository = repositories.NewCurrencyRepository()
	currencyHeatmapService = services.NewHeatmapService[models.Currency](currencyRepository)
	currencyHeatmapDrawerService = services.NewHeatmapDrawerService[models.Currency](currencyHeatmapService)

	http.Handle("/heatmap", http.HandlerFunc(heatmapHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func heatmapHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	currencyHeatmapDrawerService.Draw(w)
}
