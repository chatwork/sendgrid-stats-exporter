package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"os"
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
			[]string{"user_name", "category"},
			nil,
		),
		bounceDrops: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "bounce_drops"),
			"bounce_drops",
			[]string{"user_name", "category"},
			nil,
		),
		bounces: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "bounces"),
			"bounces",
			[]string{"user_name", "category"},
			nil,
		),
		clicks: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "clicks"),
			"clicks",
			[]string{"user_name", "category"},
			nil,
		),
		deferred: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "deferred"),
			"deferred",
			[]string{"user_name", "category"},
			nil,
		),
		delivered: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "delivered"),
			"delivered",
			[]string{"user_name", "category"},
			nil,
		),
		invalidEmails: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "invalid_emails"),
			"invalid_emails",
			[]string{"user_name", "category"},
			nil,
		),
		opens: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "opens"),
			"opens",
			[]string{"user_name", "category"},
			nil,
		),
		processed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "processed"),
			"processed",
			[]string{"user_name", "category"},
			nil,
		),
		requests: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "requests"),
			"requests",
			[]string{"user_name", "category"},
			nil,
		),
		spamReportDrops: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "spam_report_drops"),
			"spam_report_drops",
			[]string{"user_name", "category"},
			nil,
		),
		spamReports: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "spam_reports"),
			"spam_reports",
			[]string{"user_name", "category"},
			nil,
		),
		uniqueClicks: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unique_clicks"),
			"unique_clicks",
			[]string{"user_name", "category"},
			nil,
		),
		uniqueOpens: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unique_opens"),
			"unique_opens",
			[]string{"user_name", "category"},
			nil,
		),
		unsubscribeDrops: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unsubscribe_drops"),
			"unsubscribe_drops",
			[]string{"user_name", "category"},
			nil,
		),
		unsubscribes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "unsubscribes"),
			"unsubscribes",
			[]string{"user_name", "category"},
			nil,
		),
	}
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	today := time.Now()
	statistics, err := collectByDate(today)
	if err != nil {
		log.Error(err)
		return
	}

	userName := os.Getenv("SENDGRID_USER_NAME")
	category := os.Getenv("SENDGRID_CATEGORY")

	for _, stats := range statistics[0].Stats {
		ch <- prometheus.MustNewConstMetric(
			c.blocks,
			prometheus.GaugeValue,
			float64(stats.Metrics.Blocks),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.bounceDrops,
			prometheus.GaugeValue,
			float64(stats.Metrics.BounceDrops),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.bounces,
			prometheus.GaugeValue,
			float64(stats.Metrics.Bounces),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.clicks,
			prometheus.GaugeValue,
			float64(stats.Metrics.Clicks),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.deferred,
			prometheus.GaugeValue,
			float64(stats.Metrics.Deferred),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.delivered,
			prometheus.GaugeValue,
			float64(stats.Metrics.Delivered),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.invalidEmails,
			prometheus.GaugeValue,
			float64(stats.Metrics.InvalidEmails),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.opens,
			prometheus.GaugeValue,
			float64(stats.Metrics.Opens),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.processed,
			prometheus.GaugeValue,
			float64(stats.Metrics.Processed),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.requests,
			prometheus.GaugeValue,
			float64(stats.Metrics.Requests),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.spamReportDrops,
			prometheus.GaugeValue,
			float64(stats.Metrics.SpamReportDrops),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.spamReports,
			prometheus.GaugeValue,
			float64(stats.Metrics.SpamReports),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.uniqueClicks,
			prometheus.GaugeValue,
			float64(stats.Metrics.UniqueClicks),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.uniqueOpens,
			prometheus.GaugeValue,
			float64(stats.Metrics.UniqueOpens),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.unsubscribeDrops,
			prometheus.GaugeValue,
			float64(stats.Metrics.UnsubscribeDrops),
			userName,
			category,
		)
		ch <- prometheus.MustNewConstMetric(
			c.unsubscribes,
			prometheus.GaugeValue,
			float64(stats.Metrics.Unsubscribes),
			userName,
			category,
		)
	}

}
