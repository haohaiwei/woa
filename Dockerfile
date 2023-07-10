FROM golang:1.13
WORKDIR /go/src/github.com/haohaiwei/woa/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/webhook/webhook.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/haohaiwei/woa/app .
ENTRYPOINT ["./app"]
