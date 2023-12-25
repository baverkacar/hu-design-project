# Go build stage
FROM golang:1.18-alpine AS builder

# Çalışma dizinini ayarla
WORKDIR /app

# Bağımlılıkları kopyala ve indir
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Kaynak kodunu kopyala
COPY *.go ./

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
