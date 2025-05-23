# 🛒 Online Shop API

Online Shop API adalah RESTful API yang dibangun dengan **Golang**, menggunakan **Gin** sebagai web framework dan **GORM** sebagai ORM. API ini menangani berbagai fitur dasar e-commerce seperti manajemen user, produk, kategori, cart, order, dan pembayaran.

---

## 📸 Preview Project

Fitur yang tersedia:
- 👤 Autentikasi (Register/Login/JWT)
- 🛍️ Manajemen Produk dan Kategori
- 🧺 Cart & Checkout
- 💳 Pembayaran (COD, VA, Credit, PayLater)
- 📦 Order & Status
- 🔍 Search produk dengan keyword
- 🛠️ Admin Panel Endpoint (untuk pengelolaan data)

---

## ⚙️ Configuration

### Environment Variables

Buat file `.env` di root project dengan format seperti berikut:

```

APP\_PORT=8080
DB\_HOST=localhost
DB\_PORT=3306
DB\_USER=root
DB\_PASSWORD=yourpassword
DB\_NAME=online\_shop\_api
JWT\_SECRET=yourjwtsecret

````
---

## 🚀 Menjalankan Aplikasi

### 1. Clone Repo

```bash
git clone https://github.com/rizkycahyono97/online-shop-api.git
cd online-shop-api
````

### 2. Install Dependency

```bash
go mod tidy
```

### 3. Run Aplikasi

```bash
go run main.go
```

---

## Menjalankan Dengan  Docker
```bash
docker compose build
docker compose up -d
```

---

## 🧪 API Endpoint

Berikut adalah daftar sebagian endpoint umum:

### 👤 Main

| Method | Endpoint           | Deskripsi |
| ------ | ------------------ |-----------|
| POST   | `/api/v1/` | Main API  |


### 👤 Auth

| Method | Endpoint           | Deskripsi     |
| ------ | ------------------ | ------------- |
| POST   | `/api/v1/register` | Register user |
| POST   | `/api/v1/login`    | Login user    |

### 🧍 User

| Method | Endpoint       | Deskripsi      |
| ------ | -------------- | -------------- |
| GET    | `/api/v1/user` | Ambil profile  |
| PUT    | `/api/v1/user` | Update profile |
| DELETE | `/api/v1/user` | Hapus akun     |

### 📦 Produk & Kategori

| Method | Endpoint               | Deskripsi               |
| ------ | ---------------------- | ----------------------- |
| GET    | `/api/v1/products`     | List produk             |
| POST   | `/api/v1/products`     | Tambah produk (Admin)   |
| GET    | `/api/v1/products/:id` | Detail produk           |
| PUT    | `/api/v1/products/:id` | Update produk (Admin)   |
| DELETE | `/api/v1/products/:id` | Hapus produk (Admin)    |
| GET    | `/api/v1/categories`   | List kategori           |
| POST   | `/api/v1/categories`   | Tambah kategori (Admin) |

### 🛒 Cart

| Method | Endpoint           | Deskripsi            |
| ------ | ------------------ | -------------------- |
| GET    | `/api/v1/cart`     | Ambil isi cart       |
| POST   | `/api/v1/cart`     | Tambah item ke cart  |
| DELETE | `/api/v1/cart/:id` | Hapus item dari cart |

### 🧾 Order

| Method | Endpoint             | Deskripsi                |
| ------ | -------------------- | ------------------------ |
| POST   | `/api/v1/orders`     | Checkout / Buat Order    |
| GET    | `/api/v1/orders`     | Lihat semua pesanan user |
| GET    | `/api/v1/orders/:id` | Detail pesanan tertentu  |

### 💳 Pembayaran

| Method | Endpoint                            | Deskripsi                                  |
| ------ | ----------------------------------- | ------------------------------------------ |
| POST   | `/api/v1/payments`                  | Buat pembayaran                            |
| GET    | `/api/v1/user/payments`             | Lihat semua payment user login             |
| GET    | `/api/v1/payments/:user_id`         | Lihat payment berdasarkan user\_id (Admin) |
| PUT    | `/api/v1/payments/:order_id/status` | Update status payment (Admin)              |

### 🔍 Search

| Method | Endpoint         | Deskripsi              |
| ------ | ---------------- | ---------------------- |
| GET    | `/api/v1/search` | Cari produk by keyword |

---

## 🛠️ Teknologi dan Library

* [Golang](https://golang.org/)
* [Gin](https://github.com/gin-gonic/gin) - Web Framework
* [GORM](https://gorm.io/) - ORM untuk MySQL
* [JWT](https://github.com/golang-jwt/jwt) - Autentikasi token
* [MySQL](https://www.mysql.com/) - Relational Database
* [Validator](https://github.com/go-playground/validator) - Validasi data
* [godotenv](https://github.com/joho/godotenv) - Load konfigurasi dari .env
* [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Hashing password

---

## 📁 Struktur Folder

```
.
├── config/
├── controllers/
├── database/
├── services/
├── repositories/
├── middlewares/
├── models/
├── helpers/
├── routes/
├── main.go
└── go.mod
```

---

## 👤 Role dan Hak Akses

| Role     | Hak Akses                                   |
| -------- | ------------------------------------------- |
| Customer | Melihat produk, transaksi, cart, pembayaran |
| Admin    | Kelola produk, kategori, pembayaran, user   |

---

## 🧩 Fitur Mendatang

* 🔔 Notifikasi untuk status pesanan/pembayaran
* 📊 Dashboard Admin sederhana
* 📧 Email invoice / notifikasi
* 📁 Upload gambar ke Cloud Storage (e.g. Cloudinary/S3)

---

## Testing

Untuk menguji API Contact Form, Anda dapat menggunakan file koleksi Postman yang telah disediakan, import di Postman anda.

```
test/postman/online-shop-api.postman_collection.json
test/postman/online-shop-api.postman_environment.json
```

---

## 🤝 Kontribusi

Pull Request dan kontribusi sangat diterima! Silakan fork, buat branch baru dan kirim PR.

---
