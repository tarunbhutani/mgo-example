# Build stage
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules and dependencies first to optimize caching
COPY go.mod go.sum ./

# Download necessary Go modules
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app
RUN go build -o main .

# Final stage (minimal image)
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go app binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Run the Go app
CMD ["./main"]
