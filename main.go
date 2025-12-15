package main

import (
	"os"
	"time"

	"github.com/laudenlaruto/gluetun-exporter/pkg/gluetun"
	"github.com/laudenlaruto/gluetun-exporter/pkg/promexporter"
	"github.com/qdm12/log"
)

func main() {
	logger := log.New(log.SetLevel(log.LevelInfo), log.SetComponent("gluetun-exporter"))

	// Get the EXPORTER_INTERVAL environment variable
	exporterInterval := os.Getenv("EXPORTER_INTERVAL")
	if exporterInterval == "" {
		exporterInterval = "60" // Default to 60 seconds if not provided
	}

	// Convert the interval to an integer
	interval, err := time.ParseDuration(exporterInterval + "s")
	if err != nil {
		logger.Errorf("Invalid EXPORTER_INTERVAL value: %v", err)
		os.Exit(1)
	}

	// Start the Prometheus exporter server in a background goroutine
	go func() {
		logger.Info("Starting prometheus exporter...")
		promexporter.Serve()
	}()

	// Start the metric collection loop
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	controlServer := gluetun.New()

	for range ticker.C {
		// Collect metrics from the control server
		logger.Info("Updating Metrics...")
		controlServer.Collect()
		logger.Info("Updated Metrics")
	}
}
