# Build stage
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

# Run stage
FROM alpine:latest

RUN apk add --no-cache openssh-keygen

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/data ./data

# Generate SSH host key at build time
RUN mkdir -p .ssh && ssh-keygen -t ed25519 -f .ssh/id_ed25519 -N ""

EXPOSE 22097

CMD ["./app"]
