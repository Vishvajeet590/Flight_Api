FROM golang:1.22.5-alpine

# Install necessary build tools
RUN apk add --no-cache git make gcc libc-dev

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .
# Copy .env file
COPY .env .env

# Build the application
RUN make build

# Expose the port the app runs on
EXPOSE 8080

# Run the application using make
CMD ["make", "run"]