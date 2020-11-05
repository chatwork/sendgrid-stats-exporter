package main

import "github.com/prometheus/client_golang/prometheus"

type Collector struct {
	blocks           *prometheus.Desc
	bounceDrops      *prometheus.Desc
	bounces          *prometheus.Desc
	clicks           *prometheus.Desc
	deferred         *prometheus.Desc
	delivered        *prometheus.Desc
	invalidEmails    *prometheus.Desc
	opens            *prometheus.Desc
	processed        *prometheus.Desc
	requests         *prometheus.Desc
	spamReportDrops  *prometheus.Desc
	spamReports      *prometheus.Desc
	uniqueClicks     *prometheus.Desc
	uniqueOpens      *prometheus.Desc
	unsubscribeDrops *prometheus.Desc
	unsubscribes     *prometheus.Desc
}

func collector() *Collector {
	return &Collector{
		blocks: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "blocks"),
			"blocks",
			[]string{"type", "name"},
			nil,
		),
		bounceDrops: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "bounce_drops"),
			"bounce_drops",
			[]string{"type", "name"},
			nil,
		),
		bounces: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "bounces"),
			"bounces",
			[]string{"type", "name"},
			nil,
		),
		clicks: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "clicks"),
			"clicks",
			[]string{"type", "name"},
			nil,
		),
		deferred: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "deferred"),
			"deferred",
			[]string{"type", "name"},
			nil,
		),
		delivered: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "delivered"),
			"delivered",
			[]string{"type", "name"},
			nil,
		),
		invalidEmails: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "invalid_emails"),
			"invalid_emails",
			[]string{"type", "name"},
			nil,
		),
		opens: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "opens"),
			"opens",
			[]string{"type", "name"},
			nil,
		),
		processed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "processed"),
			"processed",
			[]string{"type", "name"},
			nil,
		),
		requests: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "requests"),
			"requests",
			[]string{"type", "name"},
			nil,
		),
		spamReportDrops: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "spam_report_drops"),
			"spam_report_drops",
			[]string{"type", "name"},
			nil,
		),
		spamReports: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "spam_reports"),
			"spam_reports",
			[]string{"type", "name"},
			nil,
		),
		uniqueClicks: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unique_clicks"),
			"unique_clicks",
			[]string{"type", "name"},
			nil,
		),
		uniqueOpens: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unique_opens"),
			"unique_opens",
			[]string{"type", "name"},
			nil,
		),
		unsubscribeDrops: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unsubscribe_drops"),
			"unsubscribe_drops",
			[]string{"type", "name"},
			nil,
		),
		unsubscribes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unsubscribes"),
			"unsubscribes",
			[]string{"type", "name"},
			nil,
		),
	}
}
