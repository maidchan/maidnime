# Frontend Deployment Guide - Railway

## 📋 Persiapan

Pastikan Anda sudah:
- [x] Push code ke GitHub repository
- [x] Memiliki akun Railway
- [x] Backend sudah deploy dan berjalan

## 🚀 Deploy ke Railway

### 1. Buat Service Baru di Railway

1. Login ke [Railway.app](https://railway.app)
2. Pilih project atau buat project baru
3. Click **New Service** → **GitHub Repo**
4. Pilih repository `maidnime`

### 2. Konfigurasi Service Frontend

#### **A. Set Root Directory**
Karena ini monorepo, Railway perlu tahu di mana folder frontend:

1. Buka service frontend → **Settings**
2. Scroll ke **Root Directory**
3. Set ke: `packages/frontend`
4. Click **Save**

#### **B. Set Environment Variables**

Di tab **Variables**, tambahkan:

```
VITE_API_BASE_URL=https://your-backend-url.railway.app
```

Ganti `your-backend-url.railway.app` dengan URL backend Railway Anda.

**PENTING:** Railway akan otomatis set variable `PORT`, jangan di-override!

#### **C. Deploy Settings**

Railway akan otomatis detect Dockerfile. Pastikan:
- Builder: **Dockerfile** ✅
- Dockerfile Path: `Dockerfile` (sudah default)

### 3. Deploy

1. Click **Deploy**
2. Tunggu build process selesai (2-3 menit)
3. Railway akan generate URL otomatis (misal: `https://maidnime-frontend.up.railway.app`)

## 🔍 Troubleshooting

### ❌ Frontend tidak muncul / 502 Error

**Penyebab:**
- Container tidak serve file static
- Port tidak sesuai
- Build gagal

**Solusi:**
1. Cek **Logs** di Railway dashboard
2. Pastikan ada log: `nginx: [notice] start worker processes`
3. Pastikan PORT environment variable terdeteksi

### ❌ 404 saat refresh page (selain homepage)

**Penyebab:** SPA routing tidak fallback ke index.html

**Solusi:** Sudah di-handle di `nginx.conf` dengan:
```nginx
try_files $uri $uri/ /index.html;
```

### ❌ API call gagal / CORS error

**Penyebab:**
- Backend tidak allow origin dari frontend domain
- `VITE_API_BASE_URL` salah

**Solusi:**
1. Cek `VITE_API_BASE_URL` di Railway Variables
2. Pastikan backend CORS allow frontend domain
3. Rebuil frontend setelah update variable

### ❌ Assets tidak load (404 on JS/CSS)

**Penyebab:** Base path tidak sesuai

**Solusi:**
Cek `vite.config.ts`, pastikan:
```typescript
export default defineConfig({
  base: '/', // untuk Railway
  // ...
})
```

## 🎯 Verifikasi Deployment Berhasil

Setelah deploy, test:

1. ✅ **Homepage load** → buka URL Railway
2. ✅ **Navigation works** → click menu/link
3. ✅ **Direct URL works** → buka URL langsung (misal `/search`)
4. ✅ **API calls work** → cek Network tab di DevTools
5. ✅ **Refresh works** → refresh di halaman selain homepage

## 🔄 Update Deployment

Setelah push ke GitHub:
1. Railway otomatis detect changes
2. Auto-rebuild dan deploy
3. Zero-downtime deployment

## 🆚 Alternatif: Gunakan Caddy (Opsional)

Jika nginx bermasalah, gunakan Caddy yang lebih simpel:

1. Rename `Dockerfile` ke `Dockerfile.nginx` (backup)
2. Rename `Dockerfile.caddy` ke `Dockerfile`
3. Redeploy di Railway

Caddy lebih mudah dan otomatis handle HTTPS serta compression.

## 📊 Monitoring

Check logs secara real-time:
```bash
# Di Railway Dashboard
Tab "Deployments" → Click deployment → "View Logs"
```

Log yang bagus terlihat seperti:
```
✓ Built in Xms
Creating production build...
nginx: [notice] starting nginx
```

## 🎨 Custom Domain (Opsional)

1. Buka **Settings** → **Domains**
2. Click **Add Domain**
3. Masukkan domain Anda
4. Set DNS record sesuai instruksi Railway
5. Tunggu propagasi (5-30 menit)

## 💡 Tips Optimasi

1. **Enable Build Cache:** Railway otomatis cache `node_modules`
2. **Minimize bundle size:** Sudah optimized di Vite build
3. **Gzip enabled:** Sudah di nginx.conf
4. **Asset caching:** Static files cache 1 tahun

---

**Butuh bantuan?** Cek Railway Logs untuk error message detail.
