# --- Stage 1: The Builder ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy the Go module files first. This is a caching optimization.
# As long as go.mod/go.sum don't change, Docker won't re-download dependencies.
COPY go.mod go.sum ./

# Download all the dependencies.
RUN go mod download

# --- THIS IS THE FIX ---
# Copy the ENTIRE source tree (including subdirectories like handlers/ and models/)
# into the container.
COPY . .

# Build the Go application.
RUN CGO_ENABLED=0 GOOS=linux go build -o ab-testing-backend .


# --- Stage 2: The Final Image ---
FROM alpine:latest

WORKDIR /root/

# Copy ONLY the compiled application from the 'builder' stage.
COPY --from=builder /app/ab-testing-backend .

EXPOSE 8080

CMD ["./ab-testing-backend"]