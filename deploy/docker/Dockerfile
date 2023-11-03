# Build the application from source
FROM golang:1.19.3 AS build-binary

# Copy all files
COPY ./ /

# Download dependencies
RUN go mod download

# Build binary
RUN GOOS=linux go build -o /bin/api /cmd/cv-api.go

# Deploy application binary into Debian image
FROM debian:bookworm-slim

# Set workdir
WORKDIR /

# Copy application binary
COPY --from=build-binary /bin/api /api

# Default port is 3000, so expose it
EXPOSE 3000

ENTRYPOINT [ "/api" ]