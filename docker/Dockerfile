FROM golang:1.17-alpine
WORKDIR /app
COPY . .
RUN go mod download && go build -o api-sync-service cmd/main.go
CMD ["./api-sync-service"]
