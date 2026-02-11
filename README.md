# Digital Wallet Service

Layanan dompet digital berbasis REST API yang dibangun dengan Go, menggunakan Gin framework dan PostgreSQL database.

## 📋 Fitur

* ✅ Manajemen user dan wallet
* 💰 Cek saldo wallet
* 💸 Penarikan dana (withdraw funds)
* 📊 Pencatatan transaksi otomatis
* 🔒 Clean Architecture (Controller → Service → Repository)

## 🛠️ Tech Stack

* **Language**: Go 1.24
* **Framework**: Gin Web Framework
* **Database**: PostgreSQL
* **Driver**: pgx/v5
* **Environment**: godotenv

## 📁 Struktur Proyek

```
digital-wallet-svc/
├── cmd/
│   └── main.go                 # Entry point aplikasi
├── internal/
│   ├── app/
│   │   └── user/
│   │       ├── controller/     # HTTP handlers
│   │       ├── services/       # Business logic
│   │       ├── implementations/# Repository implementations
│   │       ├── repositories/   # Data access interfaces
│   │       ├── models/         # Data structures
│   │       └── route.go        # Route registration
│   └── http/
│       └── server.go           # HTTP server setup
├── pkg/
│   └── database/               # Database connection package
├── scripts/
│   ├── user.sql               # Database schema & seed data
│   └── execute_db.go          # Migration script
├── go.mod
└── go.sum
```

## 🚀 Instalasi

### Prerequisites

* Go 1.24 atau lebih tinggi
* PostgreSQL 12 atau lebih tinggi
* Git

### Langkah-langkah

1. **Clone atau ekstrak project**

   ```bash
   unzip digital-wallet-svc.zip
   cd digital-wallet-svc
   ```
2. **Install dependencies**

   ```bash
   go mod download
   ```
3. **Setup database**
   Buat database PostgreSQL:

   ```sql
   CREATE DATABASE digital_wallet;
   ```
4. **Konfigurasi environment**
   Buat file `.env` di root folder:

   ```env
   # Database Configuration
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASS=your_password
   DB_NAME=digital_wallet
   DB_PORT=5432
   DB_SSL=disable

   # Server Configuration
   PORT=:8080
   ```
5. **Jalankan migration database**

   ```bash
   cd scripts
   go run execute_db.go
   ```

   Script ini akan membuat:

   * Table `users`
   * Table `wallets`
   * Table `transactions`
   * Data sample (2 users, 2 wallets, 3 transaksi)
6. **Jalankan aplikasi**

   ```bash
   go run cmd/main.go
   ```

   Server akan berjalan di `http://localhost:8080`

## 📡 API Endpoints

### 1. Cek Saldo Wallet

**GET**`/user/balance_wallet`

**Query Parameters:**

```
id: user ID (required)
wallet_id: wallet ID (required)
```

**Contoh Request:**

```bash
curl "http://localhost:8080/user/balance_wallet?id=1&wallet_id=1001"
```

**Response:**

```json
{
  "balance": 1000.00
}
```

### 2. Penarikan Dana

**GET**`/user/withdraw_funds`

**Query Parameters:**

```
id: user ID (required)
wallet_id: wallet ID (required)
withdrawal_amount: jumlah penarikan (required)
```

**Contoh Request:**

```bash
curl "http://localhost:8080/user/withdraw_funds?id=1&wallet_id=1001&withdrawal_amount=100"
```

**Response Success:**

```json
{
  "message": "Withdrawal successful"
}
```

**Response Error:**

```json
{
  "error": "Insufficient balance"
}
```

## 🗃️ Database Schema

### Users Table

```sql
users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
)
```

### Wallets Table

```sql
wallets (
    wallet_id INT PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users(id),
    balance DECIMAL(15,2) NOT NULL DEFAULT 0
)
```

### Transactions Table

```sql
transactions (
    transaction_id SERIAL PRIMARY KEY,
    wallet_id INT REFERENCES wallets(wallet_id),
    amount DECIMAL(15,2) NOT NULL,
    types VARCHAR(20) NOT NULL,  -- 'withdraw' atau 'deposit'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
```

## 🧪 Testing API

Gunakan tools seperti:

* **cURL** (command line)
* **Postman** (GUI)
* **Insomnia** (GUI)

Contoh testing dengan cURL:

```bash
# Test balance wallet
curl "http://localhost:8080/user/balance_wallet?id=1&wallet_id=1001"

# Test withdraw
curl "http://localhost:8080/user/withdraw_funds?id=1&wallet_id=1001&withdrawal_amount=50"
```

## 🏗️ Arsitektur

Project ini menggunakan **Clean Architecture**:

1. **Controller Layer** (`controller/`)
   * Menerima HTTP request
   * Validasi input
   * Mengembalikan HTTP response
2. **Service Layer** (`services/`)
   * Business logic
   * Validasi aturan bisnis
   * Koordinasi antar repository
3. **Repository Layer** (`implementations/`, `repositories/`)
   * Data access logic
   * Query ke database
   * Interface untuk testability

## 📝 Development

### Menambah Endpoint Baru

1. Tambahkan model di `models/user_model.go`
2. Tambahkan method di interface `repositories/user_repo.go`
3. Implementasikan di `implementations/user_impl.go`
4. Tambahkan business logic di `services/user_svc.go`
5. Buat handler di `controller/user_ctrl.go`
6. Register route di `controller/user_ctrl.go` → `RegisterRoutes()`

### Build untuk Production

```bash
# Build binary
go build -o digital-wallet-svc cmd/main.go

# Run binary
./digital-wallet-svc
```

## 🐛 Troubleshooting

### Database connection failed

* Pastikan PostgreSQL sudah running
* Cek kredensial di file `.env`
* Pastikan database sudah dibuat
* Cek firewall settings

### Port already in use

```bash
# Cari process yang menggunakan port 8080
lsof -i :8080

# Kill process (Linux/Mac)
kill -9 <PID>
```

### Migration gagal

* Pastikan file `scripts/user.sql` ada
* Cek permission user database
* Pastikan table belum ada (drop jika perlu)

## 📦 Dependencies

Main dependencies yang digunakan:

* `github.com/gin-gonic/gin` - Web framework
* `github.com/jackc/pgx/v5` - PostgreSQL driver
* `github.com/joho/godotenv` - Environment variable loader

Lihat `go.mod` untuk dependency lengkap.

## 📄 License

Project ini dibuat untuk keperluan pembelajaran dan development.

## 👥 Contributing

Untuk development:

1. Fork project ini
2. Buat branch baru (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Buat Pull Request

## 📞 Support

Jika ada pertanyaan atau issues:

* Buat issue di repository
* Email: georgeonsent10@gmail.com

---

**Happy Coding! 🚀**
