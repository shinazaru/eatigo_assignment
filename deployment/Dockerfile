FROM golang:1.16 as builder

WORKDIR /go

COPY . /go
RUN cd /go/api \
 && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

WORKDIR /root/
COPY --from=builder /go/api/main .
CMD ["./main"]