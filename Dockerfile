# Use the official Go image as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your Gin application will listen on
EXPOSE 8080

# Command to run your Go application
CMD ["./main"]
