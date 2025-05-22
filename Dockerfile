# Stage 1
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o online-shop-api .

# Stage 2
FROM alpine:latest

RUN apk add --no-cache tzdata

ENV TZ=Asia/Jakarta

RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

RUN addgroup -g 1001 binarygroup
RUN adduser -D -u 1001 -G binarygroup userapp

COPY --from=builder --chown=userapp:binarygroup /app/online-shop-api .

USER userapp

EXPOSE 8080

CMD ["./online-shop-api"]