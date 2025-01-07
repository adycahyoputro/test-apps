# pull image dari docker hub
FROM golang:alpine AS builder

# buat direktori/folder baru
WORKDIR /build

# copy semua file dari folder backend di local ke dalam folder baru
COPY . .

# download library/dependency golang
RUN go mod download

# build project di folder baru dan beri nama binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /binary

# pull image golang yang lebih kecil ukurannya
FROM alpine:3.21

# buat folder baru lagi untuk menyimpan hasil build sebelumnya
WORKDIR /app

# copy file binary dan .env ke dalam folder baru
COPY --from=builder /binary /binary
COPY --from=builder /build/.env .

# exposep port untuk bisa diakses dari local
EXPOSE 8082

# untuk menjalankan file binary
CMD [ "/binary" ]