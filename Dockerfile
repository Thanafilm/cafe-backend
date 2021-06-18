FROM golang:alpine as builder

ENV GO111MODULE=on

RUN apk update
WORKDIR /go/src/cafe-backend
COPY . .
RUN go get -d -v ./...
RUN go get -u github.com/acoshift/goreload
CMD ["goreload", "main.go"]
EXPOSE 3000
