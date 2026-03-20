# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod file
COPY go.mod ./

# (Optional) Download dependencies if you have any
RUN go mod download

# Copy the source code and templates
COPY main.go ./
COPY index.html ./

# Build the application
RUN go build -o word-it main.go

# Stage 2: Create a minimal image to run the app
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary and templates from the builder stage
COPY --from=builder /app/word-it .
COPY --from=builder /app/index.html .

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["./word-it"]
