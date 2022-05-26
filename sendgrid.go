package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	endpoint = "https://api.sendgrid.com/v3/stats"
)

type Metrics struct {
	Blocks           int64 `json:"blocks,omitempty"`
	BounceDrops      int64 `json:"bounce_drops,omitempty"`
	Bounces          int64 `json:"bounces,omitempty"`
	Clicks           int64 `json:"clicks,omitempty"`
	Deferred         int64 `json:"deferred,omitempty"`
	Delivered        int64 `json:"delivered,omitempty"`
	InvalidEmails    int64 `json:"invalid_emails,omitempty"`
	Opens            int64 `json:"opens,omitempty"`
	Processed        int64 `json:"processed,omitempty"`
	Requests         int64 `json:"requests,omitempty"`
	SpamReportDrops  int64 `json:"spam_report_drops,omitempty"`
	SpamReports      int64 `json:"spam_reports,omitempty"`
	UniqueClicks     int64 `json:"unique_clicks,omitempty"`
	UniqueOpens      int64 `json:"unique_opens,omitempty"`
	UnsubscribeDrops int64 `json:"unsubscribe_drops,omitempty"`
	Unsubscribes     int64 `json:"unsubscribes,omitempty"`
}

type Stat struct {
	Metrics *Metrics `json:"metrics,omitempty"`
}

type Statistics struct {
	Date  string  `json:"date,omitempty"`
	Stats []*Stat `json:"stats,omitempty"`
}

func collectByDate(timeStart time.Time, timeEnd time.Time) ([]*Statistics, error) {
	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	layout := "2006-01-02"
	dateStart := timeStart.Format(layout)
	dateEnd := timeEnd.Format(layout)

	query := url.Values{}
	query.Set("start_date", dateStart)
	query.Set("end_date", dateEnd)
	if *accumulatedMetrics {
		query.Set("aggregated_by", "month")
	} else {
		query.Set("aggregated_by", "day")
	}
	parsedURL.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *sendGridAPIKey))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var reader io.Reader = res.Body
	reader = io.TeeReader(reader, os.Stdout)

	switch res.StatusCode {
	case http.StatusTooManyRequests:
		return nil, fmt.Errorf("API rate limit exceeded")
	case http.StatusOK:
		var stats []*Statistics
		if err := json.NewDecoder(reader).Decode(&stats); err != nil {
			return nil, err
		}

		return stats, nil
	default:
		return nil, fmt.Errorf("status code = %d, response = %s", res.StatusCode, res.Body)
	}
}
