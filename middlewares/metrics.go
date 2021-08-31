package middlewares

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

func TimingMetricMiddleware(timer kitprometheus.Gauge) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			start := time.Now()
			defer calculateExecutionTime(start, timer)
			return next(ctx, request)

		}
	}
}

func calculateExecutionTime(start time.Time, timer kitprometheus.Gauge) {
	stop := time.Now()
	executionTime := stop.Sub(start).Milliseconds()
	lvs := []string{"method", "count", "error", "false"}
	timer.With(lvs...).Set(float64(executionTime))
}
