FROM golang:1.16 as builder
WORKDIR /go
COPY . /go
CMD [ "go", "run", "/go/api/main" ]