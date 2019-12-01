FROM golang:1.13 as builder

WORKDIR /app

COPY . /app

RUN go get github.com/yanzay/tbot/v2

RUN CGO_ENABLED=0 go build -v -o statusoli .

FROM alpine:latest

COPY --from=builder /app/statusoli /statusoli

ENTRYPOINT ["/statusoli"]
