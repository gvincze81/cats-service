FROM golang:2.18-alpine

RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make && \
    apk add git

WORKDIR /cat-webapp

COPY . .
WORKDIR /cat-webapp/cats-service

RUN go mod download
RUN go build -o /cat-api

EXPOSE 8080
CMD ["/cat-api"]
