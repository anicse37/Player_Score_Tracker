# Use official Go image
FROM golang:1.22.2

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Run tests by default
CMD ["go", "test", "-v","./..."]
