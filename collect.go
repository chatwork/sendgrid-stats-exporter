package main

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	logger log.Logger

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

func collector(logger log.Logger) *Collector {
	return &Collector{
		logger: logger,

		blocks: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "blocks"),
			"blocks",
			[]string{"user_name"},
			nil,
		),
		bounceDrops: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "bounce_drops"),
			"bounce_drops",
			[]string{"user_name"},
			nil,
		),
		bounces: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "bounces"),
			"bounces",
			[]string{"user_name"},
			nil,
		),
		clicks: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "clicks"),
			"clicks",
			[]string{"user_name"},
			nil,
		),
		deferred: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "deferred"),
			"deferred",
			[]string{"user_name"},
			nil,
		),
		delivered: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "delivered"),
			"delivered",
			[]string{"user_name"},
			nil,
		),
		invalidEmails: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "invalid_emails"),
			"invalid_emails",
			[]string{"user_name"},
			nil,
		),
		opens: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "opens"),
			"opens",
			[]string{"user_name"},
			nil,
		),
		processed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "processed"),
			"processed",
			[]string{"user_name"},
			nil,
		),
		requests: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "requests"),
			"requests",
			[]string{"user_name"},
			nil,
		),
		spamReportDrops: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "spam_report_drops"),
			"spam_report_drops",
			[]string{"user_name"},
			nil,
		),
		spamReports: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "spam_reports"),
			"spam_reports",
			[]string{"user_name"},
			nil,
		),
		uniqueClicks: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unique_clicks"),
			"unique_clicks",
			[]string{"user_name"},
			nil,
		),
		uniqueOpens: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unique_opens"),
			"unique_opens",
			[]string{"user_name"},
			nil,
		),
		unsubscribeDrops: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unsubscribe_drops"),
			"unsubscribe_drops",
			[]string{"user_name"},
			nil,
		),
		unsubscribes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unsubscribes"),
			"unsubscribes",
			[]string{"user_name"},
			nil,
		),
	}
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	today := time.Now()

	statistics, err := collectByDate(today)
	if err != nil {
		level.Error(c.logger).Log(err)

		return
	}

	for _, stats := range statistics[0].Stats {
		ch <- prometheus.MustNewConstMetric(
			c.blocks,
			prometheus.GaugeValue,
			float64(stats.Metrics.Blocks),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.bounceDrops,
			prometheus.GaugeValue,
			float64(stats.Metrics.BounceDrops),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.bounces,
			prometheus.GaugeValue,
			float64(stats.Metrics.Bounces),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.clicks,
			prometheus.GaugeValue,
			float64(stats.Metrics.Clicks),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.deferred,
			prometheus.GaugeValue,
			float64(stats.Metrics.Deferred),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.delivered,
			prometheus.GaugeValue,
			float64(stats.Metrics.Delivered),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.invalidEmails,
			prometheus.GaugeValue,
			float64(stats.Metrics.InvalidEmails),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.opens,
			prometheus.GaugeValue,
			float64(stats.Metrics.Opens),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.processed,
			prometheus.GaugeValue,
			float64(stats.Metrics.Processed),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.requests,
			prometheus.GaugeValue,
			float64(stats.Metrics.Requests),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.spamReportDrops,
			prometheus.GaugeValue,
			float64(stats.Metrics.SpamReportDrops),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.spamReports,
			prometheus.GaugeValue,
			float64(stats.Metrics.SpamReports),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.uniqueClicks,
			prometheus.GaugeValue,
			float64(stats.Metrics.UniqueClicks),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.uniqueOpens,
			prometheus.GaugeValue,
			float64(stats.Metrics.UniqueOpens),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.unsubscribeDrops,
			prometheus.GaugeValue,
			float64(stats.Metrics.UnsubscribeDrops),
			*sendGridUserName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.unsubscribes,
			prometheus.GaugeValue,
			float64(stats.Metrics.Unsubscribes),
			*sendGridUserName,
		)
	}
}
