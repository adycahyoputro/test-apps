FROM golang:alpine
WORKDIR /test-apps
COPY . /test-apps/
RUN go mod tidy
RUN go build -o binary
EXPOSE 8081
CMD [ "/test-apps/binary" ]