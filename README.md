# sendgrid-stats-exporter

## Overview

(TBW)


    +---------------------------+                          +------------+                        +--------------+
    |  SendGrid Stats API (v3)  |---(collect /v3/stats)--->|  exporter  |<---(scrape /metrics)---|  Prometheus  |
    +---------------------------+                          +------------+                        +--------------+

## Usage

```
$ make
$ ./exporter
$ SENDGRID_API_KEY="secret" SENDGRID_USER_NAME="username" SENDGRID_CATEGORY="" ./exporter 
```

```
$ curl localhost:2112/-/healthy
$ curl localhost:2112/metrics
```

### Running with Docker

(TBW)

 - docker run
 
#### Running with `docker-compose`

```
$ cp .env.example .env
$ vi .env
$ docker-compose up -d
```

You can check the metrics by accessing Prometheus ([http://127.0.0.1:9200]()).

## Building

### Building locally

```
$ make
```

### Building with Docker

```
$ docker build -t sendgrid-stats-exporter .
```
 
## Configuration

```

```

### Exporter

You can specify a user name as environment variable to identify metrics for multiple users, as well as categories. 

Name     | Description | Default
---------|-------------|----
`SENDGRID_API_KEY` | API key for calling stats API (v3) | `""`
`SENDGRID_USER_NAME` | (Optional) Label for metrics | `""`
`SENDGRID_CATEGORY` | (not implemented) | `""`


## Endpoints

Name     | Description
---------|-------------
`/metrics` | get metrics
`/-/healthy` | for health check

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