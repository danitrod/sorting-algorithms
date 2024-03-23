FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY internal internal

RUN go build -ldflags="-s -w" -o sort cmd/main.go

FROM ubuntu:22.04

WORKDIR /app

COPY --from=builder /app/sort /app/sort
COPY data /app/data
COPY scripts/experiment.sh /app/scripts/experiment.sh

RUN chmod +x /app/sort
RUN chmod +x /app/scripts/experiment.sh

ENTRYPOINT ["scripts/experiment.sh"]
