# 🚀 Go GitHub Actions

![Go CI](https://github.com/SiNopaal/go-github-actions/actions/workflows/go-ci.yml/badge.svg)

Sebuah project sederhana yang dibuat untuk mendemonstrasikan bagaimana cara setup **CI/CD (Continuous Integration & Continuous Deployment)** menggunakan **Go (Golang)** dan **GitHub Actions**.

> Cocok untuk belajar atau jadi template awal project Golang kamu yang lebih besar 🚀

---

## 📦 Teknologi yang Digunakan

- 🧠 **Go (Golang)** — Bahasa backend yang cepat dan efisien
- ⚙️ **GitHub Actions** — Untuk otomatisasi CI/CD (build, test)
- 🧪 **go test** — Untuk unit testing sederhana
- 💡 **Modular dengan `go.mod`**

---

## 🗂️ Struktur Project

```
go-github-actions/
├── main.go              # File utama
├── main_test.go         # Unit test sederhana
├── go.mod               # Deklarasi module dan dependensi
├── go.sum               # Checksum keamanan module
└── .github/
    └── workflows/
        └── go-ci.yml    # File GitHub Actions untuk CI
```

---

## 🧪 Cara Menjalankan Lokal

### 🔧 Build & Run

```bash
go mod tidy      # pastikan semua dependensi terpasang
go run main.go   # menjalankan aplikasi
```

### ✅ Test

```bash
go test ./...    # menjalankan semua unit test
```

---

## 🔄 CI/CD dengan GitHub Actions

Setiap kali ada push ke branch `main` atau pull request baru, GitHub akan menjalankan workflow otomatis yang:

1. Checkout kode dari repo
2. Setup environment Golang
3. Jalankan `go mod tidy` untuk mengatur dependensi
4. Jalankan semua unit test dengan `go test ./...`
5. Build project dengan `go build`

📂 Lokasi file workflow: `.github/workflows/go-ci.yml`

---

## 🖼️ Contoh Output

```bash
> go run main.go
Hello from GitHub Actions CI/CD!
```

```bash
> go test ./...
ok  	github.com/SiNopaal/go-github-actions	0.001s
```

---

## 🧑‍💻 Author

Made with ❤️ by [Naufal Maulana Izzuddin](https://github.com/SiNopaal)  
📍 Patukrejomulyo, Mirit, Kebumen  
🎓 Telkom University Purwokerto  
🛠️ Passionate in Go, Web Dev, and DevOps

---

## 📌 Ingin Pakai untuk Project Kamu?

Silakan fork project ini, ubah nama module (`go.mod`), dan kembangkan sesuai kebutuhanmu!  
Butuh bantuan? Feel free to reach out via GitHub or WA 📲
