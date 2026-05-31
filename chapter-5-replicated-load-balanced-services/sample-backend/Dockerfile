FROM golang:1-alpine AS builder

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/app /app

USER 1000:1000

EXPOSE 8080

ENTRYPOINT [ "/app" ]
