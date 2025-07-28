# Gunakan base image Golang
FROM golang:1.21

# Set direktori kerja di dalam container
WORKDIR /app

# Salin go.mod dan go.sum terlebih dahulu
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Salin semua file project
COPY . .

# Build binary
RUN go build -o out ./cmd/main.go

# Jalankan binary
CMD ["./out"]
