FROM golang:3.21-alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./binary
EXPOSE 8082

FROM alpine:edge
WORKDIR /apps
COPY --from=builder /build/binary ./binary
CMD [ "/apps/binary" ]