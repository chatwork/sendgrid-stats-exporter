package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"time"
)

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

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	today := time.Now()
	metrics, err := collectByDate(today, "secret")
	if err != nil {
		log.Error(err)
		return
	}

	for _, stats := range metrics[0].Stats {
		ch <- prometheus.MustNewConstMetric(
			c.blocks,
			prometheus.GaugeValue,
			float64(stats.Metrics.Blocks),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.bounceDrops,
			prometheus.GaugeValue,
			float64(stats.Metrics.BounceDrops),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.bounces,
			prometheus.GaugeValue,
			float64(stats.Metrics.Bounces),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.clicks,
			prometheus.GaugeValue,
			float64(stats.Metrics.Clicks),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.deferred,
			prometheus.GaugeValue,
			float64(stats.Metrics.Deferred),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.delivered,
			prometheus.GaugeValue,
			float64(stats.Metrics.Delivered),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.invalidEmails,
			prometheus.GaugeValue,
			float64(stats.Metrics.InvalidEmails),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.opens,
			prometheus.GaugeValue,
			float64(stats.Metrics.Opens),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.processed,
			prometheus.GaugeValue,
			float64(stats.Metrics.Processed),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.requests,
			prometheus.GaugeValue,
			float64(stats.Metrics.Requests),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.spamReportDrops,
			prometheus.GaugeValue,
			float64(stats.Metrics.SpamReportDrops),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.spamReports,
			prometheus.GaugeValue,
			float64(stats.Metrics.SpamReports),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.uniqueClicks,
			prometheus.GaugeValue,
			float64(stats.Metrics.UniqueClicks),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.uniqueOpens,
			prometheus.GaugeValue,
			float64(stats.Metrics.UniqueOpens),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.unsubscribeDrops,
			prometheus.GaugeValue,
			float64(stats.Metrics.UnsubscribeDrops),
			stats.Type,
			stats.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.unsubscribes,
			prometheus.GaugeValue,
			float64(stats.Metrics.Unsubscribes),
			stats.Type,
			stats.Name,
		)
	}

}
