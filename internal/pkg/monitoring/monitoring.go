package monitoring

import "github.com/prometheus/client_golang/prometheus"

type Monitoring interface {
	SetupMonitoring() error
	GetSuccessHits() *prometheus.CounterVec
	GetErrorsHits() *prometheus.CounterVec
	GetRequestCounter() prometheus.Counter
	GetExecution() *prometheus.HistogramVec
}
