FROM golang:latest as builder
WORKDIR /go/src
ADD . /go/src
RUN CGO_ENABLED=0 go build . && ls -l /go/src

FROM alpine:latest
WORKDIR /go
COPY --from=builder /go/src/userquake-aggregator .

EXPOSE 8080
CMD ["./userquake-aggregator", "server", "-p", "8080"]

