FROM golang:alpine
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR $GOPATH/app
COPY . $GOPATH/app
RUN go mod download
RUN go build main.go

EXPOSE 8080
ENTRYPOINT ["./main"]
