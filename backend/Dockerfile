# Use a Golang base image
FROM golang:latest

# Install Git
RUN apt-get update && apt-get install -y git

# Set the working directory
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Download module dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Run the Go app
CMD ["./main"]