FROM golang:1.7.3 as builder
WORKDIR /go/src/github.com/daniel-bryant/hancho_sample/helloworld
RUN go get -d -v golang.org/x/net/html
COPY main.go .
RUN go get github.com/daniel-bryant/honcho_sample/service_manager/gen/go/helloworld
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/daniel-bryant/hancho_sample/helloworld/main .
CMD ["./main", "-port=50052"]
