FROM golang:1.13 as builder

WORKDIR /app

COPY . /app

RUN go get github.com/yanzay/tbot/v2

RUN go build -v -o statusoli .

FROM alpine:latest

COPY --from=builder /app/statuoli /statuoli

ENTRYPOINT ["/statusoli"]
