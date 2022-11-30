package middleware

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var successCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "Success request",
	Help: "Number of foo successfully processed.",
})

var hits = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hits",
}, []string{"status", "path"})

type MetricsMiddleware struct{}

func NewMetricsMiddleware() *HTTPMiddleware {
	return &HTTPMiddleware{}
}

func (m *MetricsMiddleware) SetDefaultMetrics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)

		hits.WithLabelValues("500", r.URL.String()).Inc()

		res := w.Result()
	})
}
