package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func newMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{
		next: next,
	}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("Pushing metrics to Prometheus")
	return s.next.FetchPrice(ctx, ticker)
}
