# Use the official Golang image as a starting point
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to enable better dependency caching
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy all files from the current directory to the container's working directory
COPY . .

# Build the application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 5000

# Default command to run the application when a container based on this image is started
CMD ["./main", "-debug=true"]
