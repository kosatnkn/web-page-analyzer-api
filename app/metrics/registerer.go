package metrics

import (
	httpMetrics "github.com/kosatnkn/web-page-analyzer-api/transport/http/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

// Register registers declared metrics.
//
// This is the central location to register metrics from different
// layers of the service.
func Register() {
	prometheus.MustRegister(httpMetrics.HTTPReqDuration)
}
