# Deployment Guide - Railway

## Masalah yang Ditemukan
Website menampilkan **Error 502 Bad Gateway** karena:
1. Railway menggunakan PORT dinamis (bukan port 80 yang hardcoded)
2. Nginx tidak dikonfigurasi untuk menggunakan environment variable PORT dari Railway

## Solusi yang Diterapkan

### 1. Update Dockerfile
- Menambahkan `gettext` package untuk `envsubst` command
- Menggunakan template nginx config yang akan di-replace dengan PORT dinamis
- Menggunakan environment variable `$PORT` dari Railway

### 2. Update nginx.conf
- Mengganti `listen 80;` dengan `listen $PORT;`
- File ini sekarang menjadi template yang akan diproses oleh `envsubst`

### 3. Menambahkan railway.json
- Konfigurasi eksplisit untuk Railway agar menggunakan Dockerfile

## Langkah Deploy ke Railway

### Opsi 1: Redeploy dari Dashboard Railway
1. Buka Railway Dashboard: https://railway.app/dashboard
2. Pilih project `maidnime-frontend-production`
3. Klik tab "Settings"
4. Scroll ke bawah dan klik "Redeploy" atau "Trigger Deploy"
5. Tunggu build selesai (biasanya 2-5 menit)

### Opsi 2: Push ke Git (Jika terhubung dengan GitHub)
```bash
cd /home/maidchan/Devs/AI/maidnime
git add packages/frontend/Dockerfile packages/frontend/nginx.conf packages/frontend/railway.json packages/frontend/.dockerignore
git commit -m "fix: Railway deployment - support dynamic PORT"
git push origin main
```

Railway akan otomatis mendeteksi perubahan dan melakukan rebuild.

### Opsi 3: Deploy Manual dengan Railway CLI
```bash
# Install Railway CLI (jika belum)
npm i -g @railway/cli

# Login
railway login

# Link ke project
cd packages/frontend
railway link

# Deploy
railway up
```

## Environment Variables yang Diperlukan

Pastikan di Railway Dashboard sudah ada environment variable:
- `VITE_API_BASE_URL` - URL backend API (contoh: https://your-backend.railway.app)

## Verifikasi Deployment

Setelah deploy selesai:
1. Cek logs di Railway Dashboard untuk memastikan tidak ada error
2. Akses https://maidnime-frontend-production.up.railway.app/
3. Website seharusnya sudah bisa diakses

## Troubleshooting

### Jika masih error 502:
1. Cek logs di Railway Dashboard
2. Pastikan build berhasil (tidak ada error saat build)
3. Pastikan PORT environment variable tersedia
4. Cek apakah service sudah running dengan status "Active"

### Jika build gagal:
1. Cek apakah semua dependencies terinstall
2. Pastikan `package.json` dan `package-lock.json` ada
3. Cek logs untuk error spesifik

### Jika website blank/kosong:
1. Cek browser console untuk error JavaScript
2. Pastikan `VITE_API_BASE_URL` sudah diset dengan benar
3. Cek apakah file di folder `dist` ter-generate dengan benar
