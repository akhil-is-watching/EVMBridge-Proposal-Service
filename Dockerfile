# Stage 1: Build the Go application
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o proposal-indexer .

# Stage 2: Create a clean container with only the built binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/proposal-indexer .

# Command to run the binary
CMD ["./proposal-indexer"]

ENTRYPOINT [ "./proposal-indexer" ]