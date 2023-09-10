# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application's source code
COPY . .

# Build the Go application
RUN go build -o ./cmd/main ./cmd/main.go

# Expose the port the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./cmd/main"]