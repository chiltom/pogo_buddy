FROM golang:1.20 AS dev
WORKDIR /app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Install air for hot reload
RUN go install github.com/cosmtrek/air@latest

# Expose port and start Air for hot reloading
EXPOSE 8080
CMD ["air"]