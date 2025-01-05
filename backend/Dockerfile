FROM golang:alpine
WORKDIR /test-apps
COPY . .
RUN go mod tidy
RUN go build -o binary
EXPOSE 8082
CMD [ "/binary" ]