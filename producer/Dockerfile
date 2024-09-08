# Golang image
FROM golang:1.22-alpine AS builder  

# Set the working directory
WORKDIR /app

# Copy files
COPY . .

# Execute commands
RUN go build main.go

# Final image
FROM alpine:3.19

# Set the working directory
WORKDIR /app

# Copy files from builder
COPY --from=builder /app/main .

# Execute
CMD ["./main"]
