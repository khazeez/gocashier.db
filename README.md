# 🧾 GoCashier API

RESTful API untuk aplikasi kasir (cashier) sederhana yang dibangun dengan **Golang**, **Gin**, dan konsep **Clean Architecture**.

API ini mendukung manajemen kategori, produk, transaksi checkout, serta laporan transaksi dengan dokumentasi otomatis menggunakan Swagger.

---

# 🚀 Tech Stack

- Golang
- Gin Web Framework
- SQL Database (MySQL/PostgreSQL)
- Swaggo Swagger
- Clean Architecture

---

# 📦 Features

✅ Category Management  
✅ Product Management  
✅ Transaction Checkout  
✅ Transaction Reports (Daily & Range)  
✅ Swagger API Documentation  
✅ Clean Architecture Structure  
✅ JSON REST API  

---

# 📂 Project Structure

```bash
gocashier.db/
│
├── api/
│ ├── handler/
│ └── router.go
│
├── internal/
│ ├── models/
│ ├── repository/
│ └── services/
│
├── docs/ # auto-generated swagger docs
├── main.go
├── go.mod
└── README.md

```

---

# ⚙️ Installation

## 1️⃣ Clone Repository

```bash
git clone https://github.com/yourusername/gocashier.git
cd gocashier

```
Run Application

```bash
go run main.go
```

Server berjalan di:

http://localhost:8080




📘 Swagger Documentation
<<<<<<< HEAD

Project ini menggunakan Swaggo untuk dokumentasi API otomatis.


Generate Swagger Docs

```bash
swag init --parseInternal
```


Akses swagger
http://localhost:8080/swagger/index.html
=======

Project ini menggunakan Swaggo untuk dokumentasi API otomatis.


Generate Swagger Docs

```bash
swag init --parseInternal
```


Akses swagger
http://localhost:8080/swagger/index.html




>>>>>>> cc424712919cb77a7dea7fcd685b7ff802e19327
