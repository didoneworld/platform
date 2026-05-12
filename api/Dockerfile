FROM golang:1.21-alpine

WORKDIR /app
COPY . /app/

RUN go mod init did-one-world-api && \
    go get github.com/google/uuid && \
    go get github.com/lestrrat-go/jwx/v2/jwa && \
    go get github.com/lestrrat-go/jwx/v2/jws && \
    go build -o server ./identities.go

FROM alpine:3.19

RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -D appuser

USER appuser
COPY --from=builder /app/server /app/

EXPOSE 8000
CMD ["/app/server"]
