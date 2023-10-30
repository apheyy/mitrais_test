# Use the official Go base image
FROM golang:1.21.3

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go build -o main

# Expose the port the application will run on
EXPOSE 8000

# Command to run the Go application
CMD ["./main"]