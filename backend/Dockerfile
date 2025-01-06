FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
# COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /binary

# FROM alpine:
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=builder /binary /binary
EXPOSE 8082
CMD [ "/binary" ]