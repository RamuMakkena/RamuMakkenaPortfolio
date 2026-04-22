FROM golang:1.25.6-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN go build -o bin/server ./cmd/server

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/bin/server .
EXPOSE 8080
CMD ["./server"]
