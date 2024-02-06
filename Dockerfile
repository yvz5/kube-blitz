# Start from a base image that includes Go runtime
FROM golang:1.21 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod .
COPY go.sum .

# Download and cache dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the controller binary
RUN CGO_ENABLED=0 GOOS=linux go build -o controller ./main.go

# Start a new stage from scratch
FROM alpine:latest  

# Set the current working directory inside the container
WORKDIR /app

# Copy the controller binary from the builder stage
COPY --from=builder /app/controller .

# Run the controller binary
CMD ["./controller"]
