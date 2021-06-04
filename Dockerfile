FROM golang:1.16.4-alpine3.13
RUN mkdir /app

RUN apk add --no-cache git

# RUN go get -v -d github.com/golang-migrate/migrate/cli \
#     && go get -v -d github.com/lib/pq

COPY . /app
WORKDIR /app



CMD ["go","run", "./server.go"]

# FROM golang:1.14.6-alpine as builder
# COPY go.mod go.sum /go/src/github.com/marcellof23/GoGraphql/
# WORKDIR /go/src/github.com/marcellof23/GoGraphql/
# RUN go mod download
# COPY . /go/src/github.com/marcellof23/GoGraphql/
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/GoGraphql github.com/marcellof23/GoGraphql

# FROM alpine
# RUN apk add --no-cache ca-certificates && update-ca-certificates
# COPY --from=builder /go/src/github.com/marcellof23/GoGraphql/build/GoGraphql /usr/bin/GoGraphql
# EXPOSE 8080 8080
# # CMD ["ping","localhost"]
# CMD ["/usr/bin/GoGraphql"]
# ENTRYPOINT ["/usr/bin/GoGraphql"]