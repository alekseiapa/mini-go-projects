# Use the official Go image as the base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Build the Go app
RUN go build -o app .

# Expose the default port for the Go app
EXPOSE 8080

# Set the entrypoint to run the Go app
ENTRYPOINT ["./app"]
