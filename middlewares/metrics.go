package middlewares

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// wrapper for prometheus timer
type PrometheusTimer struct {
	Timer kitprometheus.Gauge
}

// Construct new Prometheus Timer
func NewPrometheusTimer() *PrometheusTimer {
	timer := *kitprometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Namespace: "golock",
		Subsystem: "access",
		Name:      "request_timer",
		Help:      "execution time of request",
	}, []string{"method", "error"})

	return &PrometheusTimer{
		Timer: timer,
	}
}

// Timing middleware decorator
func (pt PrometheusTimer) TimingMetricMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			start := time.Now()
			defer calculateExecutionTime(start, pt.getTimer())
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

func (pt PrometheusTimer) getTimer() kitprometheus.Gauge {
	return pt.Timer
}
