FROM golang:1.21-alpine

WORKDIR /app
COPY wallet.go /app/

RUN go mod init did-wallet && \
    go get github.com/google/uuid && \
    go build -o wallet-server /app/wallet.go

FROM alpine:3.19

RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -D appuser

USER appuser
COPY --from=builder /app/wallet-server /app/

EXPOSE 3000
CMD ["/app/wallet-server"]
