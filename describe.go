package main

import "github.com/prometheus/client_golang/prometheus"

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.blocks
	ch <- c.bounceDrops
	ch <- c.bounces
	ch <- c.clicks
	ch <- c.deferred
	ch <- c.delivered
	ch <- c.invalidEmails
	ch <- c.opens
	ch <- c.processed
	ch <- c.requests
	ch <- c.spamReportDrops
	ch <- c.spamReports
	ch <- c.uniqueClicks
	ch <- c.uniqueOpens
	ch <- c.unsubscribeDrops
	ch <- c.unsubscribes
}
