# Stage 1
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o online-shop-api

# Stage 2
FROM alpine:latest

RUN apk add --no-cache tzdata curl ca-certificates unzip tar

ENV TZ=Asia/Jakarta

RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

RUN addgroup -g 1001 binarygroup
RUN adduser -D -u 1001 -G binarygroup userapp

# ✅ Copy binary dan direktori migrations
COPY --from=builder --chown=userapp:binarygroup /app/online-shop-api .
COPY --from=builder --chown=userapp:binarygroup /app/database/migrations /app/database/migrations

# ✅ Install golang-migrate CLI
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz  \
    -o migrate.tar.gz && \
    tar -xzf migrate.tar.gz -C /usr/local/bin && \
    rm -f migrate.tar.gz

USER userapp

EXPOSE 8080

CMD ["sh", "-c", "migrate -path=/app/database/migrations -database \"mysql://${MARIADB_USER}:${MARIADB_PASSWORD}@tcp(mariadb-online-shop:3306)/${MARIADB_DATABASE}\" up && ./online-shop-api"]