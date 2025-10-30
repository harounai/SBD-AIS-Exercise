# ======================
# Stage 1: Builder
# ======================
FROM golang:1.25 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy all source files into the container
COPY . .

# Build the Go application using the provided script
RUN chmod +x scripts/build-application.sh && ./scripts/build-application.sh


# ======================
# Stage 2: Runtime
# ======================
FROM alpine:latest

# Create app directory
WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder /app/ordersystem /app/ordersystem

# Expose the port the app will listen on
EXPOSE 3000

# Run the compiled binary
CMD ["/app/ordersystem"]