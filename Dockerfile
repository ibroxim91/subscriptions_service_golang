FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o subscription-service ./cmd/main.go

# Final stage
FROM scratch
COPY --from=builder /app/subscription-service /
COPY .env /
ENTRYPOINT ["/subscription-service"]
