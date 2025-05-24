# Use an official Go image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy everything from your local folder to the container
COPY . .

# Build the Go app
RUN go build -o site .

# Expose port 8080 (the one used in your main.go)
EXPOSE 8080

# Run the binary when the container starts
CMD ["./site"]
