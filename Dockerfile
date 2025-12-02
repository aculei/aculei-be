FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o aculei-be .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/aculei-be .
EXPOSE 8888

ENTRYPOINT ["/app/aculei-be"]