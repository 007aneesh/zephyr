package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	FunctionExecTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "function_execution_time_seconds",
			Help: "Histogram of function execution times",
		},
		[]string{"function_name"},
	)
)

func StartMetricsServer(port string) {
	prometheus.MustRegister(FunctionExecTime)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+port, nil)
}
