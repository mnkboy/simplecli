FROM golang:1.25.8 AS builder

WORKDIR /simplecli

COPY . .

RUN go build -o simplecli .

ENTRYPOINT [ "./simplecli" ]

