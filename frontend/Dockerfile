# pull image dari docker hub
FROM nginx:alpine-slim

# buat direktori/folder baru
WORKDIR /frontend-build

# copy semua file dari folder frontend di local ke dalam path Nginx
COPY . /usr/share/nginx/html

# exposep port untuk bisa diakses dari local
EXPOSE 8083