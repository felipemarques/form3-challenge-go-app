FROM golang:1.12-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
#RUN go mod tidy
COPY . .

CMD CGO_ENABLED=0 go test main.go *_test.go -v