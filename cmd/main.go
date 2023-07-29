package main

import (
	"fmt"
	"log"
	"net/http"

	svg "github.com/ajstarks/svgo"
)

const (
	Width  = 1000
	Height = 1000
)

func main() {
	http.Handle("/heatmap", http.HandlerFunc(heatmap))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func heatmap(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(Width, Height)
	s.Rect(0, 0, Width, Height, "stroke:black")
	currencyRect(s, 10, 10, 200, 100, Currency{
		Name:            "Bitcoin",
		Price:           10000,
		Last24Variation: 10,
		MarketCap:       1000000000,
	})
	s.End()
}

func currencyRect(s *svg.SVG, x, y, w, h int, currency Currency) {
	color := "green"
	if currency.Last24Variation < 0 {
		color = "red"
	}
	s.Rect(x, y, w, h, "fill:"+color+";stroke:black")
	s.Text(x+10, y+20, currency.Name, "font-size:14pt;fill:white")
	s.Text(x+10, y+40, fmt.Sprintf("%.2f", currency.Price), "font-size:14pt;fill:white")
	s.Text(x+10, y+60, fmt.Sprintf("%.2f", currency.Last24Variation), "font-size:14pt;fill:white")
}

type Currency struct {
	Name            string
	Price           float64
	Last24Variation float64
	MarketCap       float64
}
