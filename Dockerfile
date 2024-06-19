# 使用官方的golang:alpine镜像作为构建镜像
FROM golang:1.20-alpine AS builder

# 安装 git 和构建工具
RUN apk add --no-cache git build-base

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制其余的*.go文件到工作目录
COPY . .

# 构建 Go 应用程序，包含所有源文件
RUN go build -o main .

# 使用一个alpine镜像来运行应用程序
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从builder镜像复制构建好的二进制文件
COPY --from=builder /app/main .

# 确保二进制文件可执行
RUN chmod +x ./main

# 暴露应用程序端口
EXPOSE 8080

# 启动应用程序
CMD ["./main"]