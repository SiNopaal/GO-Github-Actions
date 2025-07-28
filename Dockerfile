# Gunakan image Go versi terbaru yang stabil
FROM golang:1.21-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu (untuk cache dependency)
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy semua file ke container
COPY . .

# Build binary dari main.go yang ada di folder cmd/
RUN go build -v -o out ./cmd/main.go

# Buka port 8080 (sesuai server kamu)
EXPOSE 8080

# Jalankan file hasil build
CMD ["./out"]
