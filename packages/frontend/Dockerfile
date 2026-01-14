FROM node:20-alpine AS builder

WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN npm install

# Build-time env for Vite
ARG VITE_API_BASE_URL
ENV VITE_API_BASE_URL=${VITE_API_BASE_URL}

# Copy source files and build
COPY . .
RUN npm run build

# ---- Serve stage with Caddy (for Railway) ----
FROM caddy:2.7.6-alpine

WORKDIR /app

# Caddy config & built static files
COPY Caddyfile .
COPY --from=builder /app/dist ./dist

# Railway akan inject $PORT saat runtime.
# Caddyfile sudah menggunakan :{$PORT}, jadi cukup expose port default.
EXPOSE 8080

CMD ["caddy", "run", "--config", "Caddyfile", "--adapter", "caddyfile"]
