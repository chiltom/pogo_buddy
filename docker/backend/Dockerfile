FROM golang:1.23 AS dev
WORKDIR /app

# Install CompileDaemon for development hot reloading
RUN go install github.com/githubnemo/CompileDaemon@latest

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Expose application port
EXPOSE 8080

# Start the application with CompileDaemon for hot reloading
CMD ["CompileDaemon", "--build=go build -o ./tmp/main .", "--command=./tmp/main"]