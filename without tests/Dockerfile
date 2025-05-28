# Use the official Go image as a base for building the application
FROM golang:1.22.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory.
# This allows Docker to cache the go mod download step if dependencies haven't changed.
COPY go.mod ./

# Download all Go module dependencies.
# This is a separate step to leverage Docker's build cache.
RUN go mod download

# Copy the rest of the application source code into the container.
# The '.' copies everything from the build context (your current directory)
# except what's excluded by .dockerignore.
COPY . .

# Build the Go application.
# -o /app/player-score-tracker: specifies the output binary name and path.
# ./...: tells Go to build all packages in the current module.
# CGO_ENABLED=0: Disables CGO, which makes the binary statically linked and smaller.
#                This is often good practice for Docker builds.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/player-score-tracker .


# --- Stage 2: Create a minimal production image ---
# Use a minimal base image to reduce the final image size.
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /app

# Copy only the compiled binary from the 'builder' stage into the final image
COPY --from=builder /app/player-score-tracker .

# Expose port 8080 as your Go app listens on it
EXPOSE 8080

# Command to run the application when the container starts
CMD ["./player-score-tracker"]