FROM golang:1.22 AS base

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

FROM base

COPY . .
RUN go build

CMD ["./scraper"]
