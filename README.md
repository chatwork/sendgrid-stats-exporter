# sendgrid-stats-exporter

## Overview

(TBW)


    +---------------------------+                          +------------+                        +--------------+
    |  SendGrid Stats API (v3)  |---(collect /v3/stats)--->|  exporter  |<---(scrape /metrics)---|  Prometheus  |
    +---------------------------+                          +------------+                        +--------------+

## Usage

### Using Docker

(TBW)

 - docker run
 - docker-compose up -d
 
### Configuration

You can specify a user name to identify metrics for multiple users, as well as categories. 

Name     | Description | Default
---------|-------------|----
`SENDGRID_API_KEY` | API key for calling stats API (v3) | `""`
`SENDGRID_USER_NAME` | (Optional) Label for metrics | `""`
`SENDGRID_CATEGORY` | (not implemented) | `""`


## Metrics

Name     | Description
---------|------------
blocks | dummy
bounce_drops | dummy
bounces | dummy
deferred | dummy
delivered | dummy
invalid_emails | dummy
processed | dummy
requests | dummy
spam_report_drops | dummy
spam_reports | dummy
unique_clicks | dummy
unique_opens | dummy
unsubscribe_drops | dummy
unsubscribes | dummy