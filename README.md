# 🌸 Maidnime

**Maidnime** adalah platform streaming anime modern dengan desain premium, responsif, dan cepat. Dibangun menggunakan arsitektur monorepo dengan **Vue 3** di sisi frontend dan **Go** sebagai backend engine.

![Maidnime Dashboard Mockup](https://images2.alphacoders.com/131/1313627.jpg)

## ✨ Fitur Unggulan

- 🎨 **Premium UI/UX**: Tampilan modern dengan *Glassmorphism*, *Dark Mode*, dan animasi halus menggunakan Tailwind CSS.
- 🔍 **Smart Search**: Pencarian anime secara real-time yang cepat dan akurat.
- 📱 **Fully Responsive**: Pengalaman menonton yang mulus di perangkat mobile, tablet, maupun desktop.
- 🎥 **Advanced Video Player**: Player video khusus dengan kontrol kustom, pemilih kualitas, dan pemilih server.
- 📑 **Watchlist**: Simpan anime favorit Anda ke daftar putar (tersimpan secara lokal).
- ⚡ **Go-Powered Backend**: Engine backend yang ringan dan efisien menggunakan bahasa pemrograman Go.

## 🚀 Teknologi yang Digunakan

### Frontend
- **Framework**: [Vue 3 (Composition API)](https://vuejs.org/)
- **Build Tool**: [Vite](https://vitejs.dev/)
- **Styling**: [Tailwind CSS v4](https://tailwindcss.com/)
- **Language**: [TypeScript](https://www.typescriptlang.org/)
- **Routing**: [Vue Router](https://router.vuejs.org/)

### Backend
- **Language**: [Go (Golang)](https://go.dev/)
- **Web Framework**: [Gin Gonic](https://gin-gonic.com/)
- **Data Source**: Integrated Scraper (Sanka Engine)

## 🛠️ Struktur Project

```text
maidnime/
├── packages/
│   ├── anistream/       # Frontend Vue.js
│   └── backend/         # Backend API Go
├── docker-compose.yml   # Konfigurasi Docker
└── README.md
```

## 📦 Cara Menjalankan Secara Lokal

### Prasyarat
- [Node.js](https://nodejs.org/) (v18+)
- [Go](https://go.dev/) (v1.21+)

### Langkah-langkah
1. **Clone Repository**
   ```bash
   git clone https://github.com/username/maidnime.git
   cd maidnime
   ```

2. **Jalankan Backend**
   ```bash
   npm run dev:backend
   ```
   Server backend akan berjalan di `http://localhost:8080`.

3. **Jalankan Frontend**
   Buka terminal baru di root project:
   ```bash
   npm run dev
   ```
   Aplikasi akan berjalan di `http://localhost:5173`.

## 🚢 Deployment

### Frontend (GitHub Pages)
Project ini sudah dilengkapi dengan **GitHub Actions**. Cukup push ke branch `main`, dan aplikasi akan ter-deploy otomatis.
> Jangan lupa tambahkan `VITE_API_BASE_URL` di GitHub Repository Secrets.

### Backend
Backend dapat di-deploy ke layanan seperti **Render**, **Railway**, atau **Google Cloud Run** menggunakan Dockerfile yang tersedia di folder `packages/backend`.

---

## 📄 Lisensi & Disclaimer
Project ini dibuat untuk tujuan pembelajaran. Seluruh konten anime berasal dari data publik. Kami tidak menyimpan file video di server kami.

© 2026 **Maidnime Team**. All rights reserved.
