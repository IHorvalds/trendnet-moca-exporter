FROM golang:1.23.3-alpine as build-step

WORKDIR /app

COPY go.mod ./
COPY main.go metrics.go moca-status.go ./
RUN go mod tidy
RUN go mod download
RUN go build -o ./moca-exporter

FROM alpine:latest as runner

WORKDIR /app
COPY --from=build-step /app/moca-exporter .
RUN mkdir /configs
EXPOSE 8080

CMD ["./moca-exporter", "-config", "/configs/moca.toml"]
