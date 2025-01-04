FROM golang:1.23.4
RUN apk update && apk add --no-cache git
WORKDIR /test-apps
COPY . /test-apps/
RUN go mod tidy
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o bin/backend main.go
CMD [ "/test-apps/bin/backend" ]