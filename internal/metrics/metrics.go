package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type ApiMetrics struct {
	ApiRequestCounts         *prometheus.CounterVec
	APiRequestProcessingTime *prometheus.HistogramVec
}

func NewAPIMetrics() *ApiMetrics {
	apiRequestCounts := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_request_counts",
		Help: "when an end point is accessed from the server",
	}, []string{"route"})

	requestProcessingTime := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_processing_time",
			Help:    "Time taken for request to process by each API endpoint",
			Buckets: []float64{.01, .05, .1, .25, .5, 1, 2.5, 5, 10, 30, 60, 120, 240},
		}, []string{"route"})

	return &ApiMetrics{
		ApiRequestCounts:         apiRequestCounts,
		APiRequestProcessingTime: requestProcessingTime,
	}
}

func New() *prometheus.Registry {
	appRegistry := prometheus.NewRegistry()
	appRegistry.MustRegister(collectors.NewGoCollector())
	appRegistry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	return appRegistry
}
