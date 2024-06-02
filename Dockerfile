FROM golang:1.18-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

FROM scratch

# Kaynak imajdan derlenmiş uygulama ve gerekli config dosyasını kopyala
WORKDIR /root/
COPY --from=builder /app/myapp .
COPY --from=builder /app/resources/ /root/resources/

EXPOSE 8080
CMD ["./myapp"]
