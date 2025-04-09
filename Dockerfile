FROM golang:1.23-alpine as build-stage

RUN mkdir -p /app

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o click cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/click /app/click
COPY --from=build-stage /app/config /app/config
COPY --from=build-stage /app/migrations /app/migrations

EXPOSE 9999

ENTRYPOINT [ "/app/click" ]
