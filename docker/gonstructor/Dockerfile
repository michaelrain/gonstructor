FROM golang:1.21.1 as builder

RUN mkdir -p $GOPATH/src
WORKDIR $GOPATH/src
ADD . .
ENV GO111MODULE=on

RUN go build -o /bin/app ./cmd/gonstructor

FROM ubuntu:22.04

COPY --from=builder /bin/app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


COPY config.yaml /etc/gonstructor/
COPY screens.yaml .
RUN apt-get update

EXPOSE 8080

CMD ["./app"]