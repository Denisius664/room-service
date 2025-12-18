# Build stage
FROM golang:1.25-rc-bullseye AS builder
WORKDIR /src

# Cache Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /room-service ./cmd/app

# Runtime stage
FROM gcr.io/distroless/static:nonroot
COPY --from=builder /room-service /room-service
COPY --from=builder /src/config.yaml /etc/room-service/config.yaml

EXPOSE 8080

# Default configPath env - can be overridden at runtime
ENV configPath=/etc/room-service/config.yaml

USER nonroot:nonroot
ENTRYPOINT ["/room-service"]
