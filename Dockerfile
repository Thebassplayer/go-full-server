# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# The specific image version may need to be updated over time.
FROM golang:1.22 AS builder

# Set the working directory outside $GOPATH to enable Go modules support.
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire source code from the current directory to the working directory inside the container
COPY *.go ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /full-server

EXPOSE 8080
# Run the web service on container startup.
CMD ["/full-server"]