FROM golang:1.17 as mod
ARG GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct
WORKDIR /root/myapp/

FROM mod as builder
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpserver main.go


FROM alpine:3.11.6
WORKDIR /root
COPY --from=builder /root/myapp/httpserver /httpserver
ENTRYPOINT ["/httpserver"]
