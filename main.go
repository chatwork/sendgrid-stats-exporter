package main

import (
	"context"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	namespace    = "sendgrid"
	exporterName = "sendgrid-stats-exporter"
)

const (
	stopTimeoutSecond = 10
)

var (
	gitCommit     string
	listenAddress = kingpin.Flag(
		"web.listen-address",
		"Address to listen on for web interface and telemetry.",
	).Default(":9154").Envar("LISTEN_ADDRESS").String()
	disableExporterMetrics = kingpin.Flag(
		"web.disable-exporter-metrics",
		"Exclude metrics about the exporter itself (promhttp_*, process_*, go_*).",
	).Envar("DISABLE_EXPORTER_METRICS").Bool()
	sendGridUserName = kingpin.Flag(
		"sendgrid.username",
		"Set SendGrid username",
	).Default("").Envar("SENDGRID_USER_NAME").String()
	sendGridApiKey = kingpin.Flag(
		"sendgrid.api-key",
		"Set SendGrid API key",
	).Default("secret").Envar("SENDGRID_API_KEY").String()
)

func main() {
	promlogConfig := &promlog.Config{}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	kingpin.Version(version.Info())
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	logger := promlog.New(promlogConfig)

	level.Info(logger).Log("msg", "Starting", exporterName, "version", version.Info(), gitCommit)
	level.Info(logger).Log("Build context", version.BuildContext())

	level.Info(logger).Log("msg", "Listening on", *listenAddress)

	collector := collector()
	prometheus.MustRegister(collector)
	prometheus.Unregister(prometheus.NewGoCollector())
	registry := prometheus.NewRegistry()
	if !*disableExporterMetrics {
		registry.MustRegister(
			prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
			prometheus.NewGoCollector(),
		)
	}
	registry.MustRegister(collector)

	sig := make(chan os.Signal, 1)
	signal.Notify(
		sig,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer signal.Stop(sig)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	mux.HandleFunc("/-/healthy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`OK`))
	})

	srv := &http.Server{
		Addr:    *listenAddress,
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			level.Error(logger).Log("err", err)
		}
	}()

	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), stopTimeoutSecond*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		level.Error(logger).Log("err", err)
	}
}
