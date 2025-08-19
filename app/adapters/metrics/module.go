package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/fx"
)

// EnableMetrics configures and enables Prometheus metrics collection in the application.
// It registers the provided collectors in a new Prometheus registry and sets up
// an HTTP server to expose the metrics.
//
// Parameters:
//   - collector: Variadic prometheus.Collector that will be registered for metrics collection
//
// Returns:
//   - fx.Option: An fx option that configures the metrics module
//
// The created module:
//   - Registers all provided collectors
//   - Sets up an HTTP server to expose metrics
//   - Uses "liquor-adapter-metrics" as the module identifier
func EnableMetrics(collector ...prometheus.Collector) fx.Option {
	return fx.Module("lq-adapter-metrics",
		fx.Provide(func() *prometheus.Registry {
			reg := prometheus.NewRegistry()
			for _, v := range collector {
				reg.MustRegister(v)
			}
			return reg
		}),
		fx.Invoke(newMetricsServer),
	)
}
