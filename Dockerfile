# Use the official Golang image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go app into the working directory
COPY . .

# Download necessary Go modules
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["./main"]
