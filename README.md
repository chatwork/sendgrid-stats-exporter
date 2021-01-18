# sendgrid-stats-exporter

Prometheus exporter for SendGrid daily metrics exposed by SendGrid Stats API(v3).

    +---------------------------+                          +------------+                        +--------------+
    |  SendGrid Stats API (v3)  |---(collect /v3/stats)--->|  exporter  |<---(scrape /metrics)---|  Prometheus  |
    +---------------------------+                          +------------+                        +--------------+

## Usage

```
$ make
$ ./dist/exporter --sendgrid.api-key='secret' --web.listen-address=':9154' --web.disable-exporter-metrics
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
      --sendgrid.api-key="secret"
                              [Required] Set SendGrid API key
      --sendgrid.username=""  [Optional] Set SendGrid username as a label for each metrics. This is for identifying multiple SendGrid users metrics.
      --sendgrid.location=""    [Optional] Set a zone name.(e.g. 'Asia/Tokyo') The default is UTC.
      --sendgrid.time-offset=0  [Optional] Specify the offset in second from UTC as an integer.(e.g. '32400') This needs to be set along with location.
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
blocks | The number of emails that were not allowed to be delivered by ISPs.
bounce_drops | The number of emails that were dropped because of a bounce.
bounces | The number of emails that bounced instead of being delivered.
deferred | The number of emails that temporarily could not be delivered.
delivered | The number of emails SendGrid was able to confirm were actually delivered to a recipient.
invalid_emails | The number of recipients who had malformed email addresses or whose mail provider reported the address as invalid.
processed | Requests from your website, application, or mail client via SMTP Relay or the API that SendGrid processed.
requests | The number of emails that were requested to be delivered.
spam_report_drops | The number of emails that were dropped due to a recipient previously marking your emails as spam.
spam_reports | The number of recipients who marked your email as spam.
unique_clicks | The number of unique recipients who clicked links in your emails.
unique_opens | The number of unique recipients who opened your emails.
unsubscribe_drops | The number of emails dropped due to a recipient unsubscribing from your emails.
unsubscribes | The number of recipients who unsubscribed from your emails.

### Running with Docker

```
$ docker run -d -p 9154:9154 chatwork/sendgrid-stats-exporter
```
 
#### Running with `docker-compose`

```
$ cp .env.example .env
$ vi .env
$ docker-compose up -d
```

You can check the metrics by accessing Prometheus ([http://127.0.0.1:9200]()).

#### Running with `helm`

https://github.com/chatwork/sendgrid-stats-exporter/tree/main/charts

## Building

### Building locally

```
$ make
```

### Building with Docker

```
$ docker build -t sendgrid-stats-exporter .
```
