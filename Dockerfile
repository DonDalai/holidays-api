# Use a Golang base image
FROM golang:1.22 AS build-env

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and build info
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o holidays-api .

# Use a smaller base image
FROM alpine:latest

# Set necessary environment variables
ENV GIN_MODE=release

# Copy the executable from the build stage
COPY --from=build-env /app/holidays-api /app/holidays-api

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/holidays-api"]
