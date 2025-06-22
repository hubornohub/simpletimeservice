FROM golang:1.23-alpine3.22 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o myservice ./cmd/main.go

FROM alpine:3.22.0
# Create a non-root user
RUN adduser -D appuser

WORKDIR /app
COPY --from=builder /app/myservice .

# Use non-root user
USER appuser
EXPOSE 8080
CMD ["./myservice"]