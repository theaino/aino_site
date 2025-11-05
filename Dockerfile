FROM golang:1.24.6-alpine3.22 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o aino_site ./cmd

ENV SERVER_ADDR=0.0.0.0:8080

EXPOSE 8080

CMD ["/build/aino_site"]
