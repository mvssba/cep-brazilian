FROM golang:alpine

MAINTAINER Marcos Silva<mvssba@gmail.com>

# Set the Current Working Directory inside the container
WORKDIR /app

ADD . /app/

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
