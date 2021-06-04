#!/bin/bash
FROM golang:1.14.3-alpine as build

WORKDIR /src

COPY . .

RUN go build -o testing -v cmd/testing/main.go
RUN chmod a+x testing

FROM alpine as bin


WORKDIR /src


COPY --from=build /src/testing /src/
RUN chmod a+x testing

#Command to run the executable
CMD ["./testing"]