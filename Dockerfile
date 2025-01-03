FROM golang:1.23.4
WORKDIR /test-apps
COPY . .
RUN go mod tidy
RUN go build -o binary
CMD [ "/test-apps/binary" ]