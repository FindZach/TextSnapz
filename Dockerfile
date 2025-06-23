# Use the official Go image as the base image
FROM golang:1.24.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o textsnapz ./cmd/api/main.go

# Use a minimal base image for the runtime
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/textsnapz .

# Copy templates and static files
COPY templates/ ./templates/
COPY static/ ./static/

# Create a directory for the database file
RUN mkdir -p /app/data

# Set the volume for the database file
VOLUME /app/data

# Expose the new unique port
EXPOSE 9090

# Environment variables (optional, adjust as needed)
ENV PORT=9090

# Run the application
CMD ["./textsnapz"]