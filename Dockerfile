FROM 242617/go-builder:1.0.2 AS builder

WORKDIR /build
ADD . .
RUN make build

FROM alpine:3.12
RUN apk --no-cache add ca-certificates
COPY --from=builder /build/bin/app /etc/service/app

STOPSIGNAL 15

EXPOSE 8080/tcp

ENTRYPOINT ["/etc/service/app"]
