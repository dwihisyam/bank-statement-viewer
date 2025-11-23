# Bank Statement Viewer - Backend (Go)

## Deskripsi singkat
Backend sederhana untuk take-home test: menerima upload CSV transaksi, menghitung balance, dan menampilkan transaksi bermasalah (FAILED + PENDING). Data disimpan di memori.

## Struktur proyek
cmd/app/main.go
internal/
handler/
service/
repository/
model/
utils/
pkg/response/
Dockerfile


## Endpoints
- `POST /upload` - form field `file` (multipart/form-data). Mengganti seluruh data yang ada.
- `GET /balance` - mengembalikan `{ "balance": <int64> }` (credits - debits of SUCCESS).
- `GET /issues` - mengembalikan array transaksi dengan status FAILED atau PENDING.

## Cara menjalankan (lokal)
1. Pastikan Go terinstall (>= 1.20)
2. Dari folder `backend/`:
```bash
cd backend
go mod tidy
cd cmd/app
go run main.go
