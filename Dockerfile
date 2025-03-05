FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/orbit ./cmd/main.go

# Use a small alpine image for the final container
FROM alpine:3.17

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/orbit /app/

# Set environment variables
ENV PORT=8080
ENV ENV=production

# Expose the port
EXPOSE 8080

# Run the binary
CMD ["/app/orbit"]
