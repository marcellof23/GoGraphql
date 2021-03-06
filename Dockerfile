FROM golang:1.16.4-alpine3.13
RUN mkdir /app

RUN apk update --no-cache \
  && apk add --no-cache \
    git \
    gcc \
    musl-dev

WORKDIR /app
COPY . /app

RUN go mod tidy
RUN go mod vendor
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


CMD ["go","run", "./server.go"]

