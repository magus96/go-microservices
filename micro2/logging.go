package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceFetcher
}

func newLoggingService(serv PriceFetcher) PriceFetcher {
	return &loggingService{
		next: serv,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":  time.Since(begin),
			"err":   err,
			"price": price,
		}).Info("fetchPrice")
	}(time.Now())

	price, err = s.next.FetchPrice(ctx, ticker)
	return price, err
}
