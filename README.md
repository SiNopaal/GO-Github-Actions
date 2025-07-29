# 🚀 Go GitHub Actions API
![Go](https://img.shields.io/badge/Go-1.24-blue)
![Docker](https://img.shields.io/badge/Docker-Ready-green)
![Railway](https://img.shields.io/badge/Deployed-Railway-%230B0)

API sederhana berbasis Golang dengan CI/CD GitHub Actions dan auto deploy ke Railway. Proyek ini dirancang sebagai latihan dan portofolio DevOps & Backend Go sederhana.

## 🌐 Live Demo

✅ URL: https://railway.com/invite/QjZ6Pl7AP3L

## 🔧 Fitur

- `GET /hello` — Balikan JSON: `{ "message": "Hello, World!" }`
- CI/CD otomatis dengan GitHub Actions
- Build menggunakan Docker
- Auto deploy ke Railway saat push ke `main`
- Struktur clean dan siap dikembangkan

## 🧱 Stack yang Digunakan

- Golang `v1.24`
- Docker
- GitHub Actions (CI/CD)
- Railway (Hosting/Deploy)

## 📦 Struktur Folder

```
.
├── cmd
│   └── main.go        # Entry point aplikasi
├── internal
│   └── handler.go     # Handler endpoint /hello
├── go.mod
├── go.sum
├── Dockerfile         # Build Docker image
├── .github/workflows
│   └── deploy.yml     # GitHub Actions deploy ke Railway
└── README.md
```

## ▶️ Menjalankan Secara Lokal

```bash
go mod tidy
go run ./cmd/main.go
```

Buka: `http://localhost:8080/hello`

## 🐳 Jalankan dengan Docker

```bash
docker build -t go-github-actions .
docker run -p 8080:8080 go-github-actions
```

## 🔁 CI/CD Otomatis

GitHub Actions akan otomatis:
- Build project saat push ke `main`
- Deploy ke Railway dengan `RAILWAY_TOKEN`

## 🙋‍♂️ Kontributor

Made with ❤️ by [@SiNopaal](https://github.com/SiNopaal)
