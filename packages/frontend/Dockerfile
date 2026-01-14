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

# ---- Serve stage (Node + serve) ----
FROM node:20-alpine

WORKDIR /app

# Copy only built assets
COPY --from=builder /app/dist ./dist

# Install static file server
RUN npm install -g serve

# Railway biasanya pakai PORT=3000, tapi hormati env PORT
ENV PORT=3000
EXPOSE 3000

# Jalankan server static
CMD ["sh", "-c", "serve -s dist -l ${PORT:-3000}"]
