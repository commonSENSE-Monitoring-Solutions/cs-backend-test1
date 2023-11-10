
# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory to the app directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 for the app to listen on
EXPOSE 8081

# Run the app
CMD ["./main"]
