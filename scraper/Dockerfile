# --- BUILD BASE -----
FROM golang:1.22-alpine AS build-base

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

# --------------------
# --- BUILD BINARY ---
FROM build-base AS build-bin

COPY . .
RUN go build

# --------------------
# --- RUN BASE -------
FROM alpine:3 AS run-base

RUN apk update && apk add --no-cache chromium=~126

# --------------------
# --- RUN ------------
FROM run-base
COPY --from=build-bin /usr/src/app/scraper /usr/local/bin/scrapelunch
ENTRYPOINT [ "scrapelunch" ]
