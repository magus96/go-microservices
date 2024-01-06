package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFether struct{}

func (s *priceFether) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMock = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 2_000.0,
	"ATK": 4_32.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {

	price, ok := priceMock[ticker]
	if !ok {
		return price, fmt.Errorf("The given ticker (%s) is not supported", ticker)
	}

	return price, nil
}
