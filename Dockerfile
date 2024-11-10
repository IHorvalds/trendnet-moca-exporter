FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod ./
COPY main.go metrics.go moca-status.go ./
RUN mkdir /configs
RUN go mod tidy
RUN go mod download
RUN go build -o ./moca-exporter
EXPOSE 8080

CMD ["./moca-exporter", "-config", "/configs/moca.toml"]