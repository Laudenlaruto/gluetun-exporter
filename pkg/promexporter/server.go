package promexporter

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/qdm12/log"
)

var (
	logger = log.New(log.SetLevel(log.LevelInfo), log.SetComponent("gluetun-exporter"))
)

func Serve() {
	port := os.Getenv("EXPORTER_PORT")
	if port == "" {
		port = "8001"
	}

	logger.Info("Registering metrics...")
	RegisterControlServerMetrics()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Errorf("Failed to start prometheus exporter: %v", err)
	}
}
