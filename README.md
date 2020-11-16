# sendgrid-stats-exporter

## Overview

(TBW)


    +---------------------------+                          +------------+                        +--------------+
    |  SendGrid Stats API (v3)  |---(collect /v3/stats)--->|  exporter  |<---(scrape /metrics)---|  Prometheus  |
    +---------------------------+                          +------------+                        +--------------+

## Usage

```
$ make
$ ./exporter --sendgrid.api-key='secret' --web.listen-address=':9154' --web.disable-exporter-metrics
```

```
$ curl localhost:9154/-/healthy
$ curl localhost:9154/metrics
```

```
$ ./exporter -h
usage: exporter [<flags>]

Flags:
  -h, --help                  Show context-sensitive help (also try --help-long and --help-man).
      --web.listen-address=":9154"
                              Address to listen on for web interface and telemetry.
      --web.disable-exporter-metrics
                              Exclude metrics about the exporter itself (promhttp_*, process_*, go_*).
      --sendgrid.username=""  Set SendGrid username
      --sendgrid.api-key="secret"
                              Set SendGrid API key
      --log.level=info        Only log messages with the given severity or above. One of: [debug, info, warn, error]
      --log.format=logfmt     Output format of log messages. One of: [logfmt, json]
      --version               Show application version.
```

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
