FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
# COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /binary
COPY . .

FROM alpine:3.21
# FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /app
COPY --from=builder /binary .env ./
EXPOSE 8082
CMD [ "/binary" ]