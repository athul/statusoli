FROM alpine:latest

COPY /build/statusoli /statusoli

RUN chmod +x /statusoli

ENTRYPOINT ["/statusoli"]
