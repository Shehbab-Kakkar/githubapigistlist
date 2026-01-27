# ----------- Build Stage -----------
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod ./
RUN go mod download

# Copy all source files
COPY . .

# Build statically linked Linux binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# ----------- Final Image -----------
FROM gcr.io/distroless/base-debian12

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server /server

# Expose port 8080
EXPOSE 8080

# Use non-root user
USER nonroot:nonroot

# Start server
ENTRYPOINT ["/server"]
