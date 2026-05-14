# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o figureshelf-backend ./cmd/server/main.go


# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/figureshelf-backend .

EXPOSE 8080

CMD ["./figureshelf-backend"]