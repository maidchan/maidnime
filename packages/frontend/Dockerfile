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

# Set default PORT if not provided (Railway uses 3000 by default)
ENV PORT=3000

# Expose port (Railway will inject $PORT env var, typically 3000)
EXPOSE 3000

# Create startup script that substitutes PORT in Caddyfile
RUN echo '#!/bin/sh' > /app/start.sh && \
    echo 'set -e' >> /app/start.sh && \
    echo 'PORT=${PORT:-3000}' >> /app/start.sh && \
    echo 'echo "Starting Caddy on port $PORT"' >> /app/start.sh && \
    echo 'if [ ! -d "/app/dist" ]; then' >> /app/start.sh && \
    echo '  echo "ERROR: /app/dist directory not found!"' >> /app/start.sh && \
    echo '  exit 1' >> /app/start.sh && \
    echo 'fi' >> /app/start.sh && \
    echo 'sed "s/:3000/:$PORT/g" /app/Caddyfile > /tmp/Caddyfile' >> /app/start.sh && \
    echo 'caddy run --config /tmp/Caddyfile --adapter caddyfile' >> /app/start.sh && \
    chmod +x /app/start.sh

CMD ["/app/start.sh"]
