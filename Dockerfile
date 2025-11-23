# ========== 1. Build Stage ==========
FROM node:20-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy package files first (layer caching)
COPY package.json package-lock.json* yarn.lock* ./

# Install dependencies
RUN npm install

# Copy all source code
COPY . .

# Build Next.js (output in .next)
RUN npm run build

# ========== 2. Runtime Stage ==========
FROM node:20-alpine

# Set working directory inside container
WORKDIR /app

# Copy node_modules from builder
COPY --from=builder /app/node_modules ./node_modules

# Copy built Next.js app
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json

# Expose frontend port
EXPOSE 3000

# Start production server
CMD ["npm", "run", "start"]
