FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o waste-management-service ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/waste-management-service .

EXPOSE 8080

CMD ["./waste-management-service"]