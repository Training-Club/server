# Use a specific version of golang for a reproducible build environment
FROM golang:1.22 AS builder

# Enable Go modules and set the GOPROXY environment variable to ensure faster and more reliable builds
ENV GO111MODULE=on \
    GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

# Copy the go mod and sum files first to leverage Docker cache layering,
# so these steps are only re-executed when go.mod or go.sum change
COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the application source code
COPY . .
COPY /bin ./bin

# Compile the binary with flags to reduce size and disable CGO
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main .

# Use a distroless base image for security and minimal size
FROM gcr.io/distroless/base-debian10

WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the config.yaml file from the project into the container
# Assuming /bin/config.yaml is relative to the project root in the build context
COPY --from=builder /app/bin/config.yaml ./bin

# Expose port 8080 (ensure your application is configured to use this port)
EXPOSE 8080

# Command to run
CMD ["./main"]
