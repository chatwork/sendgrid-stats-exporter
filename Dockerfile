FROM golang:1.15.4-alpine3.12 as build

COPY ./ /go/src/github.com/chatwork/sendgrid-stats-exporter
WORKDIR /go/src/github.com/chatwork/sendgrid-stats-exporter

RUN go mod download \
#    && go test ./... \
    && CGO_ENABLED=0 GOOS=linux go build -o /bin/exporter

FROM alpine:3.12

RUN apk --no-cache add ca-certificates \
     && addgroup exporter \
     && adduser -S -G exporter exporter
USER exporter
COPY --from=build /bin/exporter /bin/exporter
ENV LISTEN_PORT=2112
EXPOSE 2112
ENTRYPOINT [ "/bin/exporter" ]