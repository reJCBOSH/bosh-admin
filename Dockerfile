# 使用官方Go运行时作为基础镜像
FROM golang:1.24.2-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go mod和sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用小型基础镜像
FROM alpine:latest

# 安装ca证书（解决https请求问题）
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从builder阶段复制构建好的二进制文件
COPY --from=builder /app/main .

# 复制配置文件
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/ip2region_v4.xdb .

# 创建静态文件目录
RUN mkdir -p static

# 暴露端口
EXPOSE 8089

# 运行应用
CMD ["./main"]