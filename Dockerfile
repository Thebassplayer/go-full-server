# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# The specific image version may need to be updated over time.
FROM golang:1.22.0darwin AS builder

# Set the working directory outside $GOPATH to enable Go modules support.
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire source code from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN go build -o ./go-full-server ./main.go

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# The specific image version may need to be updated over time.
FROM alpine:latest AS runner

WORKDIR /app

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/go-full-server .

# Run the web service on container startup.
CMD ["./go-full-server"]