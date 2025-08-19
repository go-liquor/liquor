package metrics

import (
	"fmt"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/go-liquor/liquor/v3/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func newMetricsServer(cfg *config.Config, reg *prometheus.Registry) {
	port := cfg.GetInt("metrics.port")
	path := cfg.GetString("metrics.path")
	if port == 0 {
		port = 8181
	}
	if path == "" {
		path = "/metrics"
	}

	var server *gin.Engine
	if cfg.GetBool(config.AppDebug) {
		gin.SetMode(gin.DebugMode)
		server = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		server = gin.New()
	}

	server.GET(path, ginprom.PromHandler(promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		Registry: reg,
	})))
	go server.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
