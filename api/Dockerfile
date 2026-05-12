FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY api/ /app/api/

RUN go mod init did-one-world-api && \
    go get github.com/google/uuid && \
    go get github.com/lestrrat-go/jwx/v2/jwa && \
    go get github.com/lestrrat-go/jwx/v2/jws && \
    go build -o /vc-server ./api/credentials/

FROM alpine:3.19

RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser

USER appuser
WORKDIR /app
COPY --from=builder /vc-server /app/

EXPOSE 8001 8004

CMD ["/app/vc-server"]
