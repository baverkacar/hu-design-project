# Go build stage
FROM golang:1.22.1-alpine AS builder

# Çalışma dizinini ayarla
WORKDIR /app

# Go modül dosyalarını kopyala
COPY go.mod ./
COPY go.sum ./

# Bağımlılıkları indir
RUN go mod download

# Kaynak kodunu kopyala
COPY . ./

# Uygulamayı build et
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

# Çalışma dizinini ayarla
WORKDIR /root/

# Build edilmiş binary'i kopyala
COPY --from=builder /app/main .

# Uygulamayı çalıştır
CMD ["./main"]

