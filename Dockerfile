FROM golang:alpine AS builder
WORKDIR /src/
COPY helloserver.go go.* /src/
RUN go build -o /bin/helloserver
RUN chmod a+x /bin/helloserver

FROM alpine

COPY --from=builder /bin/helloserver /bin/helloserver

CMD ["/bin/helloserver"]

