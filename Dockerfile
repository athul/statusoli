FROM golang:1.13 as builder

WORKDIR /app

COPY . /app

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN CGO_ENABLED=0 go build -v -o statusoli .

FROM alpine:latest

COPY --from=builder /app/statuoli /statuoli

ENTRYPOINT ["/statusoli"]
