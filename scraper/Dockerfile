# --- BASE -----------
FROM golang:1.22-alpine AS base

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
# --------------------
# --- BUILD BINARY ---
FROM base AS build-bin

COPY . .
RUN go build
# --------------------
# --------------------
FROM alpine:3

RUN apk add chromium

COPY --from=build-bin /usr/src/app/scraper /usr/local/bin/scrapelunch
ENTRYPOINT [ "scrapelunch" ]
