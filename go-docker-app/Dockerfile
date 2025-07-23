# Use official golang image as base
FROM golang:1.24 As builder

# Set working directory
WORKDIR /app

COPY go.mod ./
RUN go mod download

# Copy the current directory to the container
COPY . .

#  Build the Go app
RUN go build -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080

# Command to run the executable
CMD ["./app"]