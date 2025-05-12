FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

# Install dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build ./main.go

# Use a small alpine image for the final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/ .
COPY --from=builder /app/internal/config ./internal/config

EXPOSE 8082 50052
CMD ["./main"]