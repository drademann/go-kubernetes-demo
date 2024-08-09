FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./greeter .

RUN go build -o greeter


FROM debian:stable-slim

COPY --from=builder /app/greeter /greeter

EXPOSE 8080
