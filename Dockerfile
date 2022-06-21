FROM golang:1.18-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.yaml .
COPY db/migrations ./migrations
COPY start.sh .
COPY wait-for.sh .
RUN apk add curl\
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz\
    && mv migrate /app/migrate

EXPOSE 8080
CMD ["/app/main"]