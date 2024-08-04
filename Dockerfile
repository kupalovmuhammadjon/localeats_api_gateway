# Stage 1: Build stage
FROM golang:1.22.5 AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Stage 2: Final stage
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .

# Copy the .env file if it's needed
COPY --from=builder /app/.env .

# Expose port 8888
EXPOSE 8888

# Command to run the executable
CMD ["./myapp"]
