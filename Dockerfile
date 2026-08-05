# --- Stage 1: Build Stage ---
FROM golang:1.25.3-alpine AS builder

# Install gcc and musl-dev required for CGO compilation
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy dependency definitions first to leverage layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build the binary with CGO enabled
# CGO_ENABLED=1 is required for go-sqlite3
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o language-learner .

# --- Stage 2: Runtime Stage ---
FROM alpine:latest

# Install runtime C library dependencies and CA certificates
RUN apk add --no-cache sqlite-libs ca-certificates

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/language-learner .
# Copy your assets/templates folder into the runtime container
COPY --from=builder /app/assets ./assets
# Copy your assets/static folder into the runtime container
COPY --from=builder /app/static ./static

# Expose port (adjust if your application uses a different port)
EXPOSE 8080

# Run the binary
CMD ["./language-learner"]