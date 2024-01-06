package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type PriceResponse struct {
	ticker string  `json:"ticker"`
	price  float64 `json:"price"`
}

type JSONAPIServer struct {
	svc PriceFetcher
}

func (s *JSONAPIServer) Run() {

}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceresponse := PriceResponse{
		price:  price,
		ticker: ticker,
	}
	return writeJSON(w, http.StatusOK, &priceresponse)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
