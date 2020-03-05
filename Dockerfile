FROM golang:1.13 as builder
WORKDIR /go/src/github.com/TV2-Bachelorproject/server
COPY . .

ENV CGO_ENABLED=0
RUN go build

FROM alpine
COPY --from=builder \
  /go/src/github.com/TV2-Bachelorproject/server/server \
  /usr/bin/server

CMD ["/usr/bin/server"]
