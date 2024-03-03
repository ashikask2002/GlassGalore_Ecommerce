# Use the official Golang image as a base
FROM golang:1.21.3-alpine3.18 AS build-stage

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project into the container
COPY ./ .

# Create a build directory
RUN mkdir -p /app/build

# Download Go modules
RUN go mod download

# Build the Go application
RUN go build -v -o /app/build/api ./cmd/api

# Use a minimal base image for the final stage
FROM gcr.io/distroless/static-debian11

# Set the working directory inside the container
WORKDIR /

# Copy the built binary from the previous stage into the final image
COPY --from=build-stage /app/build/api /api

# Copy any additional files your application may need (e.g., templates, environment files)
COPY --from=build-stage /app/template /template
COPY --from=build-stage /app/.env /

# Expose the port that your application listens on
EXPOSE 3000

# Define the command to run your application
CMD ["/api"]
