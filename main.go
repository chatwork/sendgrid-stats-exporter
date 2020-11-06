package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"log"
	"net/http"
)

const (
	namespace    = "sendgrid"
	exporterName = "sendgrid-stats-exporter"
)

const (
	port = 2112
)

func main() {
	log.Printf("Starting %s %s\n", exporterName, version.Info())
	log.Printf("Build context %s\n", version.BuildContext())

	log.Printf("Listening on %d", port)

	collector := collector()
	prometheus.MustRegister(collector)

	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
