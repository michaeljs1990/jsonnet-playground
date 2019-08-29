# build stage
FROM golang:1.12.9-buster AS builder

RUN apt update && \
    apt install git

ADD . /src

RUN cd /src && \
    go build -o jplay .

# final stage
FROM debian:buster

ENV PORT 8080

WORKDIR /app

COPY --from=builder /src/jplay .
COPY static static

ENTRYPOINT ["/app/jplay"]
