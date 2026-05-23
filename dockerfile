# =========================
# ETAPA 1: BUILD
# =========================

FROM golang:1.26-alpine AS builder

# Instalar git
RUN apk add --no-cache git

# Directorio de trabajo
WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el resto del proyecto
COPY . .

# Compilar binario
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/api

# =========================
# ETAPA 2: PRODUCCIÓN
# =========================

FROM alpine:latest

WORKDIR /app

# Copiar binario desde builder
COPY --from=builder /app/server .

# Exponer puerto
EXPOSE 8080

# Ejecutar aplicación
CMD ["./server"]